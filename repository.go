package main

import (
	"context"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository interface {
	Create(*cap.CapDownscale) error
	GetAll() ([]*cap.CapDownscale, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
// https://stackoverflow.com/questions/55564562/what-is-the-bson-syntax-for-set-in-updateone-for-the-official-mongo-go-driver
func (repository *MongoRepository) Create(downscale *cap.CapDownscale) error {
	filter := bson.D{{"accountID", downscale.AccountID}}
	update := bson.D{{"$inc", bson.D{
		{"value", -downscale.Delta},
	}}}

	upsert := true

	if _, err := repository.collection.UpdateOne(context.Background(), filter, update, &options.UpdateOptions{Upsert: &upsert} ); err != nil {
		print(err)
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
