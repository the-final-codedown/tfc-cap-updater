// shippy-cli-consignment/main.go
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/micro/go-micro"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
)

const (
	address         = "localhost:50051"
	defaultFilename = "cap.json"
)

func parseFile(file string) (*cap.CapDownscale, error) {
	var consignment *cap.CapDownscale
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("tfc.cap.cli"))
	service.Init()

	client := cap.NewCapUpdaterServiceClient("tfc.cap.updater", service.Client())

	// Contact the server and print out its response.
	/*file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}*/

	var downscale cap.CapDownscale
	downscale.AccountID = 10
	downscale.Delta = 100

	r, err := client.DownscaleCap(context.Background(), &downscale)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Accepted)
}
