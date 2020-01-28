package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto"
	"log"
	"math"
)

func (h *handler) ReadCapUpdates() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"broker:29092"},
		Topic:     "kafka-cap",
		GroupID:   "cap-reader",
		Partition: 0,
		MinBytes:  1,    // 10KB
		MaxBytes:  10e6, // 10MB
	})

	log.Println("Inside update")
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Message: %v", string(m.Value))
		log.Printf("Offset: %v", m.Offset)
		log.Printf("Key: %v", string(m.Key))
		for _, h := range m.Headers {
			log.Printf("%v - %v", h.Key, string(h.Value))
		}
		log.Printf("================")
		result := struct {
			Money               int32
			AmountSlidingWindow int32
		}{}
		if err := json.Unmarshal(m.Value, &result); err != nil {
			log.Println("Error in parsing GET cap request")
		}
		downscale := cap.CapDownscale{
			AccountID: string(m.Key),
			Value:     int32(math.Min(float64(result.Money), float64(result.AmountSlidingWindow))),
			Money:     result.Money}
		_, err = h.UpscaleCap(context.Background(), &downscale)
		if err != nil {
			log.Println(err)
		}
	}
}
