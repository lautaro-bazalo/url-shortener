package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	App      App   `json:"app"`
	Database DB    `json:"db"`
	Cache    Cache `json:"cache"`
}
type App struct {
	Port string `json:"port"`
}
type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Cache struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func ReadFromFile() (*Config, error) {
	var config *Config

	bytes, err := os.ReadFile("/app/config/config.json")

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func readEnvVars(config *Config) *Config {

	pass := os.Getenv("DB_PASSWORD")

	config.Database.Password = pass

	return config
}
