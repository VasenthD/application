package services

import (
	"application/models"
	pro "application/proto"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"


	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pro.UnimplementedChannelServiceServer
}

var (
	ChannelCollection *mongo.Collection
	Ctx               context.Context
)

func (s *Server) AddChannel(ctx context.Context, req *pro.AddChannelReq) (*pro.ViewChannelRes, error) {

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

	temp:=find.Decode(&chanInfo)
	if temp != nil {
		fmt.Println("can't find channels:  ", err)
		return nil, err
	}

	var foldersResp []*pro.FoldResp
	for _, i:= range chanInfo.Folders{
		foldResp := &pro.FoldResp{
			FoldID:   i.FoldID,
			Path:     i.Path,
			GPGKeyID: i.GPGKeyID,
		}
		foldersResp = append(foldersResp,foldResp)
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

func (s *Server) ViewChannel(ctx context.Context, req *pro.ChanReq) (*pro.ViewChannelRes, error) {

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
			Path:      i.Path,
			GPGKeyID:  i.GPGKeyID,
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
		FoldersResp: 		foldersResp,
	}

	return res, nil
}

//for pagenation
var CurrentPage = 1

func (s *Server) ListChannel(ctx context.Context, emp *pro.Empty) (*pro.ListChannelRes, error) {

	count,err := ChannelCollection.CountDocuments(context.TODO(),bson.D{})
	if err != nil {
		return nil, err
	}

	//3 records per page
	limit := int64(3)
	skip  :=  int64(CurrentPage - 1) * limit
	fOpt := options.FindOptions{
		Limit:               &limit,
		Skip:                &skip,
	}

	temp := count % limit
	pages := 0
	if temp != 0{
		pages = (int(count) / int(limit)) + 1
	} else{
		pages = (int(count) / int(limit))
	}
	
	if pages < CurrentPage{
		CurrentPage = 1
	}

	fmt.Println(count, " ",CurrentPage, " ",skip)
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
		if err != nil{
			fmt.Println("error in decoding to chaninfo: ", err)
			return nil, err
		}

		id := chanInfo.ID.Hex()
		var foldersResp []*pro.FoldResp
		for _, i := range chanInfo.Folders {
			fold := &pro.FoldResp{
				FoldID:    i.FoldID,
				Path:      i.Path,
				GPGKeyID:  i.GPGKeyID,
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
			FoldersResp: 		foldersResp,
		}

		chanList = append(chanList, chanl)
	}
	res := &pro.ListChannelRes{
		ListResp: chanList,
	}
	
	return res, nil
}