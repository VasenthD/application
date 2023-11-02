package main

import (
	pro "application/proto"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func Put(appClient pro.ApplicationServiceClient) {

	ctx := context.Background()
	clientStream, err := appClient.PutFile(ctx)
	if err != nil {
		fmt.Println("error putting into stream: ", err)
	}

	localFile, err := os.Open("/home/vasenth/download1.jpeg")
	if err != nil {
		fmt.Println("error opening local file: ", err)
	}
	stat, _ := localFile.Stat()
	fmt.Println("size: ", stat.Size())
	buffer := make([]byte, 5000)
	batchNumber := 1

	for {
		num, err := localFile.Read(buffer)
		if err == io.EOF {
			fmt.Println("reached end of file....")
			chunk := buffer[:num]
			if err := clientStream.Send(&pro.PutFileReq{
				AppID:    "6542074fbd2904ee6c40ab08",
				FoldID:   "ac1c3f60-b5f2-46f6-ad9c-20a6c66304ef",
				FileName: "newNovFlower.jpeg",
				Chunk:    chunk,
			}); err != nil {
				fmt.Println("error streaming file: ", err)
			}
			// time.Sleep(15 * time.Second)
			break
		}
		if err != nil {
			fmt.Println("error reading local file: ", err)
		}

		chunk := buffer[:num]
		if err := clientStream.Send(&pro.PutFileReq{
			AppID:        "6542074fbd2904ee6c40ab08",
			FoldID:       "ac1c3f60-b5f2-46f6-ad9c-20a6c66304ef",
			FileName:     "newNovFlower.jpeg",
			Chunk:        chunk,
			RequiredSize: 5000,
			StreamedSize: int32(num),
		}); err != nil {
			fmt.Println("error streaming file: ", err)
		}

		time.Sleep(5 * time.Second)
		fmt.Printf("Sent - batch %v - size - %v bytes\n", batchNumber, len(chunk))
		batchNumber += 1
	}
}

func Get(appClient pro.ApplicationServiceClient) {
	ctx := context.TODO()
	req := pro.GetFileReq{
		AppID:    "6542074fbd2904ee6c40ab08",
		FoldID:   "ac1c3f60-b5f2-46f6-ad9c-20a6c66304ef",
		FileName: "abcd.txt",
	}

	serverStream, err := appClient.GetFile(ctx, &req)
	if err != nil {
		log.Println("error get file: ", err)
	}

	var fileSize = 0
	newLocalFile, _ := os.Create("/home/vasenth/Documents/newLocalFile.txt")

	for {
		fmt.Println("1")
		res, err := serverStream.Recv()

		if err == io.EOF {
			log.Println("reached end of file: ", err)
			break
		}
		if err != nil {
			log.Println("Error in receiving bytes: ", err)
		}

		chunk := res.GetChunk()
		fmt.Println("4")
		fileSize += len(chunk)
		fmt.Println("Received chunk with size: ", fileSize)
		_, err1 := newLocalFile.Write(chunk)
		if err1 != nil {
			log.Fatal("error writing chunk: ", err1)
		}
		fmt.Println("5")
	}
}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error Dialing: ", err)
	}
	appClient := pro.NewApplicationServiceClient(conn)
	Put(appClient)
	// Get(appClient)
}
