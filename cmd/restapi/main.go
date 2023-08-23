package main

import (
	"flag"
	"fmt"
	"github.com/dmitry1721/authRestApi/internal/config"
	"github.com/dmitry1721/authRestApi/internal/storage/mongo"
	"go.uber.org/zap"
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

	fmt.Println(cfg)

	logger := configureLogger(cfg.Env)

	logger.Info("server started successfully!", zap.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	mongoDB, err := mongo.New("users")
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	_ = mongoDB

	// TODO: init router: http

	// TODO: run server
}

func configureLogger(env string) *zap.Logger {
	var logger *zap.Logger
	//defer logger.Sync()

	switch env {
	case envLocal:
		logger = zap.NewExample()
	case envProd:
		logger, _ = zap.NewProduction()
	}

	return logger
}
