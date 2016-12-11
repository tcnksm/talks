package main

import (
	"encoding/json"
	"log"
	"os"
)

const defaultOutput = "./golang-tokyo.json"

func main() {
	config := &Config{
		Name: "deeeet",
	}

	if err := SaveConfig(config, defaultOutput); err != nil {
		log.Fatal(err)
	}
}

func SaveConfig(config *Config, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	endcoder := json.NewEncoder(f)
	if err := endcoder.Encode(config); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Name string `json:"name"`
}
