package service

import "errors"

var (
	ErrInvalidToken = errors.New("invalid token provided")
	ErrTokenExpired = errors.New("token has expired")
)
