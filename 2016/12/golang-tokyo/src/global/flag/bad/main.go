package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

const defaultConfigPath = "./golang-tokyo.json"

var cfgPath = flag.String("config", defaultConfigPath, "")

func main() {
	flag.Parse()
	config, err := LoadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello,", config.Name)
}

func FutureStupidFunction() {
	f, err := os.Open(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
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
