package main

import (
	"encoding/json"
	"log"
	"os"
)

var configPath = "./golang-tokyo.json"

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello,", config.Name)
}

func LoadConfig() (*Config, error) {
	f, err := os.Open(configPath)
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
