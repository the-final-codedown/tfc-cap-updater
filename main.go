package main

import (
	"context"
	"fmt"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	defaultPort = ":50051"
	defaultHost = "mongodb://localhost:27017"
)

func getCaps() ([]byte, error) {
	response, err := http.Get("http://app:8081/accounts/ID/caps")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		return data, nil
	}
	/*jsonData := map[string]string{"money": "", "cap": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}*/
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

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

	cap.RegisterCapUpdaterServiceServer(s, h)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
