package handler

import (
	"encoding/json"
	"github.com/dmitry1721/authRestApi/internal/model"
	"go.uber.org/zap"
	"net/http"
)

type AuthService interface {
	GenerateJWT(*model.User) (map[string]string, error)
	UserExist(string) (*model.User, error)
	RefreshJWT(string, *model.User) (map[string]string, error)
}

type Request struct {
	Id           string `json:"id"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Response struct {
	Status       string `json:"status"`
	Error        string `json:"error,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func Auth(log *zap.Logger, s AuthService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			responseError(w, "request parameters not passed")
			log.Error("request parameters not passed")
			return
		}

		u, err := s.UserExist(id)
		if err != nil {
			log.Error(err.Error())
			responseError(w, "user not found")
			return
		}

		tokens, err := s.GenerateJWT(u)
		if err != nil {
			log.Error(err.Error())
			responseError(w, "token creation failed")
			return
		}

		responseOK(w, tokens["access_token"], tokens["refresh_token"])
		log.Info("authentication successful")
	}
}

func Refresh(log *zap.Logger, s AuthService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh_token := r.URL.Query().Get("refresh_token")
		id := r.URL.Query().Get("id")
		if refresh_token == "" && id == "" {
			msg := "request parameters not passed"
			responseError(w, msg)
			log.Error(msg)
			return
		}
		if refresh_token == "" || id == "" {
			msg := "not all required parameters passed"
			responseError(w, msg)
			log.Error(msg)
			return
		}

		u, err := s.UserExist(id)
		if err != nil {
			log.Error(err.Error())
			responseError(w, "user not found")
			return
		}

		tokens, err := s.RefreshJWT(refresh_token, u)
		if err != nil {
			log.Error(err.Error())
			responseError(w, err.Error())
			return
		}

		responseOK(w, tokens["access_token"], tokens["refresh_token"])
		log.Info("regeneration successful")
	}
}

func responseOK(w http.ResponseWriter, accessToken, refreshToken string) {
	json.NewEncoder(w).Encode(
		Response{
			Status:       StatusOK,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

func responseError(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(
		Response{
			Status: StatusError,
			Error:  msg,
		},
	)
}
