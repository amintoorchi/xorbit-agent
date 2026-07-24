package config

import (
	"encoding/json"
	"os"
)

const ConfigPath = "/etc/xorbit/config.json"

type Config struct {
	Token string `json:"token"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}