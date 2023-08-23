package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Env          string `yaml:"env"`
	DatabaseUrl  string `yaml:"database_url"`
	DatabaseName string `yaml:"database_name"`
	PrivateKey   string `yaml:"private_key"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func Load(configPath string) *Config {
	var config Config

	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
