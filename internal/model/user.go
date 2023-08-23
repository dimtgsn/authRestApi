package model

import "time"

type User struct {
	ID                  string    `bson:"_id"`
	RefreshToken        string    `bson:"refresh_token"`
	RefreshTokenExpires time.Time `bson:"refresh_token_expires"`
}
