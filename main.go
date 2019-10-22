package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	"log"
	"os"
)

const (
	port        = ":50051"
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("go.micro.api.cap"),
	)
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	capCollection := client.Database("tfc").Collection("cap")

	repository := &MongoRepository{capCollection}
	h := &handler{repository}

	// Register handlers
	cap.RegisterCapUpdaterServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
