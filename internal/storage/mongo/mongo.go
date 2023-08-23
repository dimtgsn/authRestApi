package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB *mongo.Database
}

func New(dbName string) (*Storage, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		return nil, err
	}

	if client.Ping(context.TODO(), nil) != nil {
		return nil, err
	}

	if client.Disconnect(context.TODO()) != nil {
		return nil, err
	}

	return &Storage{
		DB: client.Database(dbName),
	}, nil
}
