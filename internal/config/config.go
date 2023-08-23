package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config ...
type Config struct {
	Env         string `yaml:"env"`
	DatabaseUrl string `yaml:"database_url"`
	HTTPServer  `yaml:"http_server"`
}

// HTTPServer ...
type HTTPServer struct {
	Address     string `yaml:"address"`
	Timeout     string `yaml:"timeout"`
	IdleTimeout string `yaml:"idle_timeout"`
}

func Load(configPath string) *Config {
	var config Config

	fmt.Println(configPath)
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
