package main

import (
	"encoding/json"
	"os"
)

func LoadConfig(cfgPath string) (*Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	config := &Config{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, err
}

type Config struct {
	Name string `json:"name"`
}
