package mongo

import (
	"context"
	"github.com/dmitry1721/authRestApi/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserStorage struct {
	storage *Storage
}

func NewUserStorage(storage *Storage) *UserStorage {
	return &UserStorage{
		storage: storage,
	}
}

func (us *UserStorage) GetById(id string) (*model.User, error) {
	u := &model.User{}
	ctx := context.TODO()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objectId}}

	if err := us.storage.DB.Collection("users").FindOne(ctx, filter).Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (us *UserStorage) SaveRefreshToken(u *model.User, hashRefreshToken string) error {
	ctx := context.TODO()

	objectId, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "refresh_token", Value: hashRefreshToken},
		primitive.E{Key: "refresh_token_expires", Value: time.Now().Add(720 * time.Hour)},
	}}}

	if us.storage.DB.Collection("users").FindOneAndUpdate(ctx, filter, update).Decode(u) != nil {
		return err
	}

	return nil
}
