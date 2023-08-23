package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB *mongo.Database
}

func New(databaseUrl, databaseName string) (*Storage, error) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(databaseUrl),
	)

	if err != nil {
		return nil, err
	}

	if client.Ping(context.TODO(), nil) != nil {
		return nil, err
	}

	db := client.Database(databaseName)
	if db.CreateCollection(context.TODO(), "users") != nil {
		return nil, err
	}
	return &Storage{
		DB: client.Database(databaseName),
	}, nil
}
