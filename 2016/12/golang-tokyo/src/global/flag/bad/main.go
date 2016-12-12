package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

const defaultConfigPath = "./golang-tokyo.json"

// START OMIT
var cfgPath = flag.String("config", defaultConfigPath, "") // グローバル変数として定義

func main() {
	flag.Parse()
	config, err := LoadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello,", config.Name)
}

// END OMIT

func FutureStupidFunction() {
	f, err := os.Open(*cfgPath)
	// ...
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
