package main

import (
	"context"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type repository interface {
	Downscale(*cap.CapDownscale) error
	Upscale(*cap.CapDownscale) error
	GetAll() ([]*cap.CapDownscale, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Downscale -
// https://stackoverflow.com/questions/55564562/what-is-the-bson-syntax-for-set-in-updateone-for-the-official-mongo-go-driver
func (repository *MongoRepository) Downscale(downscale *cap.CapDownscale) error {
	filter := bson.M{"_id": downscale.AccountID}
	update := bson.D{{"$inc", bson.D{
		{"value", -downscale.Value},
		{"money", -downscale.Value},
	}}}

	if _, err := repository.collection.UpdateOne(context.Background(), filter, update); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Upscale -
// https://stackoverflow.com/questions/55564562/what-is-the-bson-syntax-for-set-in-updateone-for-the-official-mongo-go-driver
func (repository *MongoRepository) Upscale(downscale *cap.CapDownscale) error {
	filter := bson.M{"_id": downscale.AccountID}
	update := bson.D{{"$set", bson.D{
		{"value", downscale.Value},
		{"money", downscale.Money},
	}}}

	if _, err := repository.collection.UpdateOne(context.Background(), filter, update); err != nil {
		log.Println(err)
	}
	return nil
}

// GetAll -
func (repository *MongoRepository) GetAll() ([]*cap.CapDownscale, error) {
	cur, err := repository.collection.Find(context.Background(), nil, nil)
	var downscales []*cap.CapDownscale
	for cur.Next(context.Background()) {
		var downscale *cap.CapDownscale
		if err := cur.Decode(&downscale); err != nil {
			return nil, err
		}
		downscales = append(downscales, downscale)
	}
	return downscales, err
}
