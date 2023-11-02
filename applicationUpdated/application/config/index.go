package config

import (
	"application/constants"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/sftp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/ssh"
)

func ConnectDatabase() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	mongoClient := options.Client().ApplyURI(constants.ConnectionString)
	mongoConn, err := mongo.Connect(ctx, mongoClient)
	if err != nil {
		log.Println("Can't connect to mongo Client: ", err)
	}
	return mongoConn
}

func GetCollection(db string, coll string, client *mongo.Client) (collection *mongo.Collection) {
	mongoCollection := client.Database(db).Collection(coll)
	return mongoCollection
}


func SftpClient() (*sftp.Client) {
	config := &ssh.ClientConfig{
		User: "netsys",
		Auth: []ssh.AuthMethod{
			ssh.Password("Netsys@4321!"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "18.216.98.58:22", config)
	if err != nil {
		log.Println("error dialing to ssh: ", err)
	}

	client, err1 := sftp.NewClient(conn)
	if err1 != nil {
		log.Println("error connecting to sftp: ", err1)
	}
	fmt.Println("SFTP clientCreated")
	
	return client
}
