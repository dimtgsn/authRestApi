package service

import (
	"encoding/base64"
	"github.com/dmitry1721/authRestApi/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type UserStorage interface {
	GetById(string) (*model.User, error)
	SaveRefreshToken(*model.User, string) error
}

type AuthService struct {
	storage    UserStorage
	privateKey []byte
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func New(storage UserStorage, privateKey []byte) *AuthService {
	return &AuthService{
		storage:    storage,
		privateKey: privateKey,
	}
}

func (s *AuthService) UserExist(id string) (*model.User, error) {
	u, err := s.storage.GetById(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *AuthService) GenerateJWT(u *model.User) (map[string]string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * time.Minute)

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
		UserID: u.ID,
	})

	accessToken, err := at.SignedString(s.privateKey)
	if err != nil {
		return map[string]string{}, err
	}

	randomStr := getRandomString(32)

	hashRefreshToken, err := hashJWT(randomStr)
	if err != nil {
		return map[string]string{}, err
	}

	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(randomStr))

	if err = s.storage.SaveRefreshToken(u, hashRefreshToken); err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshTokenBase64,
	}, nil
}

func hashJWT(refreshToken string) (string, error) {
	hashToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	return string(hashToken), err
}

func (s *AuthService) RefreshJWT(refreshToken string, u *model.User) (map[string]string, error) {
	rf, _ := base64.StdEncoding.DecodeString(refreshToken)
	err := bcrypt.CompareHashAndPassword([]byte(u.RefreshToken), rf)
	if err != nil {
		return map[string]string{}, ErrInvalidToken
	}

	if u.RefreshTokenExpires.Before(time.Now()) {
		return map[string]string{}, ErrTokenExpired
	}

	tokens, err := s.GenerateJWT(u)
	if err != nil {
		return map[string]string{}, err
	}

	return tokens, nil
}

func getRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
