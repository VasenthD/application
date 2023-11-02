package services

import (
	"application/models"
	pro "application/proto"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/sftp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server1 struct {
	pro.UnimplementedApplicationServiceServer
}

type Server2 struct {
	pro.UnimplementedChannelServiceServer
}

var (
	ApplicationCollection *mongo.Collection
	ChannelCollection     *mongo.Collection
	SftpClient            *sftp.Client
	Ctx                   context.Context
)

func (s *Server1) Create(ctx context.Context, req *pro.CreateRequest) (*pro.CreateResponse, error) {
	time := time.Now()
	application := &models.Application{
		Id:        primitive.NewObjectID(),
		Name:      req.Name,
		ChannelId: req.ChannelId,
		CreatedAt: time.Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Format("2006-01-02 15:04:05"),
		CreatedBy: req.CreatedBy,
	}
	res, err := ApplicationCollection.InsertOne(ctx, application)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(res)
	return &pro.CreateResponse{ChannelId: req.ChannelId}, nil
}

func (s *Server1) List(ctx context.Context, req *pro.ListRequest) (*pro.ListResponse, error) {
	id, _ := primitive.ObjectIDFromHex(req.Id)
	filter := bson.M{"_id": id}
	var application *models.Application
	res, err := ApplicationCollection.Find(ctx, filter)
	if err != nil {
		log.Println("err in finding: ", err)
		return nil, err
	}
	for res.Next(ctx) {
		if err := res.Decode(&application); err != nil {
			log.Println("err in decoding: ", err)
			return nil, err
		}
	}
	stringObjectID := application.Id.Hex()
	response := &pro.ListResponse{
		Id:        stringObjectID,
		Name:      application.Name,
		ChannelId: application.ChannelId,
		CreatedBy: application.CreatedBy,
		CreatedAt: application.CreatedAt,
		UpdatedAt: application.UpdatedAt,
	}
	return response, nil
}

var CurrentPage = 1

func (s *Server1) ListAll(ctx context.Context, in *pro.Empty) (*pro.ListAllResponse, error) {
	limit := 10
	opts := options.Count().SetHint(bson.M{})
	count, err := ApplicationCollection.CountDocuments(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	temp := int(count) % int(limit)
	pages := (int(count) / int(limit))
	if temp != 0 {
		pages += 1
	}
	if pages < CurrentPage {
		CurrentPage = 1
	}
	skip := (CurrentPage - 1) * limit
	fmt.Println(skip)
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	res, err := ApplicationCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var applications []*pro.ListResponse
	for res.Next(ctx) {
		var app models.Application
		if err := res.Decode(&app); err != nil {
			log.Println("err", err)
		}
		appID := app.Id.Hex()
		temp := &pro.ListResponse{
			Id:        appID,
			Name:      app.Name,
			ChannelId: app.ChannelId,
			CreatedBy: app.CreatedBy,
			CreatedAt: app.CreatedAt,
			UpdatedAt: app.UpdatedAt,
		}
		applications = append(applications, temp)
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	res.Close(ctx)
	CurrentPage++
	response := &pro.ListAllResponse{
		ListAll: applications,
	}
	return response, nil
}

func (s *Server2) AddChannel(ctx context.Context, req *pro.AddChannelReq) (*pro.ViewChannelRes, error) {

	var folders []models.Fold
	for _, i := range req.Folders {
		fold := models.Fold{
			FoldID:   uuid.New().String(),
			Path:     i.Path,
			GPGKeyID: i.GPGKeyID,
		}
		folders = append(folders, fold)
	}

	db := &models.AddChannelReq{
		ID:                 primitive.NewObjectID(),
		Name:               req.Name,
		ChannelType:        req.ChannelType.String(),
		ServerIP:           req.ServerIP,
		AuthenticationType: req.AuthenticationType,
		Username:           req.Username,
		Password:           req.Password,
		PrivateKey:         req.PrivateKey,
		Folders:            folders,
	}

	//adding a channel
	_, err := ChannelCollection.InsertOne(ctx, db)
	if err != nil {
		fmt.Println("can't insert channel:  ", err)
		return nil, err
	}

	var chanInfo models.AddChannelReq
	query := bson.M{"_id": db.ID}
	find := ChannelCollection.FindOne(Ctx, query)

	temp := find.Decode(&chanInfo)
	if temp != nil {
		fmt.Println("can't find channels:  ", err)
		return nil, err
	}

	var foldersResp []*pro.FoldResp
	for _, i := range chanInfo.Folders {
		foldResp := &pro.FoldResp{
			FoldID:   i.FoldID,
			Path:     i.Path,
			GPGKeyID: i.GPGKeyID,
		}
		foldersResp = append(foldersResp, foldResp)
	}

	res := &pro.ViewChannelRes{
		ID:                 chanInfo.ID.Hex(),
		Name:               chanInfo.Name,
		ChannelType:        chanInfo.ChannelType,
		ServerIP:           chanInfo.ServerIP,
		AuthenticationType: chanInfo.AuthenticationType,
		Username:           chanInfo.Username,
		Password:           chanInfo.Password,
		PrivateKey:         chanInfo.PrivateKey,
		FoldersResp:        foldersResp,
	}

	return res, nil
}

func (s *Server2) ViewChannel(ctx context.Context, req *pro.ChanReq) (*pro.ViewChannelRes, error) {

	objId, err := primitive.ObjectIDFromHex(req.ChannelID)
	if err != nil {
		fmt.Println("error in converting objID: ", err)
		return nil, err
	}

	var chanInfo models.AddChannelReq
	query := bson.M{"_id": objId}
	find := ChannelCollection.FindOne(ctx, query).Decode(&chanInfo)
	if find == mongo.ErrNoDocuments {
		fmt.Println("no doc exits: ", find)
		return nil, err
	}
	if find != nil {
		fmt.Println("error in decoding chanInfo: ", find)
		return nil, err
	}

	id := req.ChannelID
	var foldersResp []*pro.FoldResp
	for _, i := range chanInfo.Folders {
		fold := &pro.FoldResp{
			FoldID:   i.FoldID,
			Path:     i.Path,
			GPGKeyID: i.GPGKeyID,
		}
		foldersResp = append(foldersResp, fold)
	}

	res := &pro.ViewChannelRes{
		ID:                 id,
		Name:               chanInfo.Name,
		ChannelType:        chanInfo.ChannelType,
		ServerIP:           chanInfo.ServerIP,
		AuthenticationType: chanInfo.AuthenticationType,
		Username:           chanInfo.Username,
		Password:           chanInfo.Password,
		PrivateKey:         chanInfo.PrivateKey,
		FoldersResp:        foldersResp,
	}

	return res, nil
}

// for pagenation
var CurrentPage2 = 1

func (s *Server2) ListChannel(ctx context.Context, emp *pro.EmptyReq) (*pro.ListChannelRes, error) {

	count, err := ChannelCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	//3 records per page
	limit := int64(3)
	skip := int64(CurrentPage-1) * limit
	fOpt := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}

	temp := count % limit
	pages := 0
	if temp != 0 {
		pages = (int(count) / int(limit)) + 1
	} else {
		pages = (int(count) / int(limit))
	}

	if pages < CurrentPage2 {
		CurrentPage2 = 1
	}

	fmt.Println(count, " ", CurrentPage, " ", skip)
	CurrentPage++

	find, err := ChannelCollection.Find(ctx, bson.M{}, &fOpt)
	if err != nil {
		fmt.Println("error in finding all docs: ", err)
		return nil, err
	}

	var chanList []*pro.ViewChannelRes
	var chanInfo models.AddChannelReq
	for find.Next(ctx) {
		err := find.Decode(&chanInfo)
		if err != nil {
			fmt.Println("error in decoding to chaninfo: ", err)
			return nil, err
		}

		id := chanInfo.ID.Hex()
		var foldersResp []*pro.FoldResp
		for _, i := range chanInfo.Folders {
			fold := &pro.FoldResp{
				FoldID:   i.FoldID,
				Path:     i.Path,
				GPGKeyID: i.GPGKeyID,
			}
			foldersResp = append(foldersResp, fold)
		}

		chanl := &pro.ViewChannelRes{
			ID:                 id,
			Name:               chanInfo.Name,
			ChannelType:        chanInfo.ChannelType,
			ServerIP:           chanInfo.ServerIP,
			AuthenticationType: chanInfo.AuthenticationType,
			Username:           chanInfo.Username,
			Password:           chanInfo.Password,
			PrivateKey:         chanInfo.PrivateKey,
			FoldersResp:        foldersResp,
		}

		chanList = append(chanList, chanl)
	}
	res := &pro.ListChannelRes{
		ListResp: chanList,
	}

	return res, nil
}

//first get into the db, get chan id, then
//find the file in sftp using req path
//receive in bytes and write


func (s *Server1) PutFile(stream pro.ApplicationService_PutFileServer) error {
	
	fmt.Println("0")
	flag := 0
	var fileSize = 0
	var newRemoteFile *sftp.File
	var chunk []byte
	
	for {
		req, err := stream.Recv()
		fmt.Println("1")
		
		// if req.StreamedSize < req.RequiredSize{
		// 	chunk = req.GetChunk()
		// 	fmt.Println("4")
		// 	fileSize += len(chunk)
		// 	fmt.Println("Received chunk with size: ", fileSize)
		// 	fmt.Println(newRemoteFile.Name())
		// 	_, err1 := newRemoteFile.Write(chunk)
		// 	if err1 != nil {
		// 		log.Fatal("error writing chunk: ", err1)
		// 	}
		// 	// log.Println("server side reached end of file: ", err)
		// 	return stream.SendAndClose(&pro.PutFileResp{
		// 		Message: "Received file size " + strconv.Itoa(fileSize),
		// 	})
		// }

		if err == io.EOF {
			chunk = req.GetChunk()
			fmt.Println("4")
			fileSize += len(chunk)
			fmt.Println("Received chunk with size: ", fileSize)
			fmt.Println(newRemoteFile.Name())
			_, err1 := newRemoteFile.Write(chunk)
			if err1 != nil {
				log.Fatal("error writing chunk: ", err1)
			}
			return stream.SendAndClose(&pro.PutFileResp{
				Message: "Received file size " + strconv.Itoa(fileSize),
			})
		}

		if err != nil {
			log.Println("Error in receiving bytes: ", err)
			return err
		}

		if flag == 0 {
			fmt.Println("2")
			newRemoteFile, _ = CreateRemoteFile(req.AppID, req.FoldID, req.FileName)
			fmt.Println("3")
			flag++
		}

		chunk = req.GetChunk()
		fmt.Println("4")
		fileSize += len(chunk)
		fmt.Println("Received chunk with size: ", fileSize)
		fmt.Println(newRemoteFile.Name())
		_, err1 := newRemoteFile.Write(chunk)
		if err1 != nil {
			log.Fatal("error writing chunk: ", err1)
		}
		fmt.Println("5")
	}

	// return stream.SendAndClose(&pro.PutFileResp{
	// 	Message: "Received file size " + strconv.Itoa(fileSize),
	// })

}

func CreateRemoteFile(appID string, foldID string, fileName string) (*sftp.File, error) {
	var ctx context.Context
	var application models.Application
	var channel models.AddChannelReq

	appObjID, _ := primitive.ObjectIDFromHex(appID)
	app_query := bson.M{"_id": appObjID}
	app := ApplicationCollection.FindOne(ctx, app_query)
	err1 := app.Decode(&application)
	if err1 != nil {
		log.Println("error decoding application: ", err1)
		return nil, err1
	}

	chanObjID, _ := primitive.ObjectIDFromHex(application.ChannelId)
	chan_query := bson.M{"_id": chanObjID}
	chanl := ChannelCollection.FindOne(ctx, chan_query)
	err2 := chanl.Decode(&channel)
	if err2 != nil {
		log.Println("error decoding channel: ", err2)
		return nil, err2
	}

	//channel found
	// get into the cghanl and loop through folders, find the folder with,req foldID
	var foldPath string

	for _, fold := range channel.Folders {
		if fold.FoldID == foldID {
			foldPath = fold.Path + "/" + fileName
		}
	}

	//sftp/testint/new.txt

	newRemoteFile, err := SftpClient.Create(foldPath)
	if err != nil {
		log.Println("error creating new remote file: ", err)
		return nil, err
	}
	return newRemoteFile, nil
}

func (s *Server1) GetFile(req *pro.GetFileReq, stream pro.ApplicationService_GetFileServer) error {

	remoteFile, _ := FindRemoteFile(req.AppID, req.FoldID, req.FileName)
	// stat, _ := remoteFile.Stat()
	buffer := make([]byte, 10)
	batchNumber := 1

	for {
		num, err := remoteFile.Read(buffer)
		if err == io.EOF {
			fmt.Println("reached end of file....")
			break
		}
		if err != nil {
			fmt.Println("error reading local file: ", err)
			return err
		}

		chunk := buffer[:num]
		if err := stream.Send(&pro.GetFileResp{
			Message: "sending chunks...",
			Chunk:   chunk,
		}); err != nil {
			fmt.Println("error streaming file: ", err)
			return err
		}

		fmt.Printf("Sent - batch %v - size - %v bytes\n", batchNumber, len(chunk))
		batchNumber += 1
	}
	return nil
}

// only folder path will be there, using that folder path,
// and the file name, we have to search in sftp and stream the file
func FindRemoteFile(appID string, foldID string, fileName string) (*sftp.File, error) {
	var ctx context.Context
	var application models.Application
	var channel models.AddChannelReq

	appObjID, _ := primitive.ObjectIDFromHex(appID)
	app_query := bson.M{"_id": appObjID}
	app := ApplicationCollection.FindOne(ctx, app_query)
	err1 := app.Decode(&application)
	if err1 != nil {
		log.Println("error decoding application: ", err1)
		return nil, err1
	}

	chanObjID, _ := primitive.ObjectIDFromHex(application.ChannelId)
	chan_query := bson.M{"_id": chanObjID}
	chanl := ChannelCollection.FindOne(ctx, chan_query)
	err2 := chanl.Decode(&channel)
	if err2 != nil {
		log.Println("error decoding channel: ", err2)
		return nil, err2
	}

	// var foldPath string
	var flag = 0
	var remoteFile *sftp.File
	var remoteFileName string
	for _, fold := range channel.Folders {
		if fold.FoldID == foldID {
			files, err := SftpClient.ReadDir(fold.Path)
			if err != nil {
				log.Fatal("Unable to list remote dir: ", err)
			}

			fmt.Println("files: ", files)

			for _, file := range files {
				fmt.Println("filename: ", file.Name())
				if file.Name() == fileName {
					flag++
					remoteFileName = fold.Path + "/" + file.Name()
					remoteFile, err = SftpClient.Open(remoteFileName)
					if err != nil {
						log.Println("error reading remote file: ", err)
						return nil, err
					}
				}
				if flag == 1 {
					break
				}
			}

			if flag == 0 {
				log.Println("remote file doesn't exist in the folder ", fold.Path, " with the name ", fileName, " ", remoteFileName)
				return nil, err
			}

		}
	}
	return remoteFile, nil
}

// func CheckFileExists(foldPath string)(bool, error) {

// }
