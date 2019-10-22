package main

import (
	"context"
	cap "github.com/the-final-codedown/tfc-cap-updater/proto/tfc/cap/updater"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(downscale *cap.CapDownscale) error
	GetAll() ([]*cap.CapDownscale, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(downscale *cap.CapDownscale) error {
	_, err := repository.collection.InsertOne(context.Background(), downscale)
	return err
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
