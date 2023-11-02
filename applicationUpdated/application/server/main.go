package main

import (
	"application/config"
	pro "application/proto"
	"application/services"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	mongoClient := config.ConnectDatabase()
	services.ApplicationCollection = config.GetCollection("testing", "Application", mongoClient)
	services.ChannelCollection = config.GetCollection("application", "channels", mongoClient)
	services.Ctx = context.Background()
	services.SftpClient = config.SftpClient()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pro.RegisterApplicationServiceServer(s, &services.Server1{})
	pro.RegisterChannelServiceServer(s, &services.Server2{})

	fmt.Println("Server listening on ", ":50051")

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
