package main

import (
	"flag"
	"github.com/dmitry1721/authRestApi/internal/config"
	"github.com/dmitry1721/authRestApi/internal/rest/handler"
	"github.com/dmitry1721/authRestApi/internal/service"
	"github.com/dmitry1721/authRestApi/internal/storage/mongo"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var configPath string

const (
	envLocal = "local"
	envProd  = "prod"
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/local.yaml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := config.Load(configPath)

	logger := configureLogger(cfg.Env)

	mongoDB, err := mongo.New(cfg.DatabaseUrl, cfg.DatabaseName)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	userStorage := mongo.NewUserStorage(mongoDB)
	authService := service.New(userStorage, []byte(cfg.PrivateKey))

	authHandleFunc := handler.Auth(logger, authService)
	refreshHandleFunc := handler.Refresh(logger, authService)

	http.HandleFunc("/auth/sign-in", authHandleFunc)
	http.HandleFunc("/auth/refresh", refreshHandleFunc)

	logger.Info("server started successfully!", zap.String("env", cfg.Env))

	srv := &http.Server{
		Addr:         cfg.Address,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("failed to start server")
	}

	logger.Error("server stopped")
}

func configureLogger(env string) *zap.Logger {
	var logger *zap.Logger

	switch env {
	case envLocal:
		logger = zap.NewExample()
	case envProd:
		logger, _ = zap.NewProduction()
	}

	return logger
}
