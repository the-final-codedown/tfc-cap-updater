package main

import (
	"context"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

const (
	defaultPort = ":50051"
	defaultHost = "mongodb://localhost:27017"
)

func main() {

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(1000),
		grpc.ConnectionTimeout(5*time.Second),
		grpc.MaxConcurrentStreams(10))

	uri := ""
	if os.Getenv("DB_HOST") != "" && os.Getenv("DB_PORT") != "" {
		uri = "mongodb://" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	} else {
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
	go func() {
		h.ReadCapUpdates()
	}()

	cap.RegisterCapUpdaterServiceServer(s, h)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
