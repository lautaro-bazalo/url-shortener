package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	App      App      `json:"app"`
	Database Database `json:"db"`
	Cache    Cache    `json:"cache"`
}
type App struct {
	Port string `json:"port"`
}
type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Cache struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func ReadFromFile() (*Config, error) {
	var config *Config
	var path string

	env := os.Getenv("environment")

	if env == "local" {
		path = "./internal/config/config.json"
	} else {
		path = "/app/config/config.json"
	}

	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return config, nil
}
