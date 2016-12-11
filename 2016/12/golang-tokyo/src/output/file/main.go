package main

import (
	"encoding/json"
	"io"
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

// START OMIT
func SaveConfig(config *Config, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	return Save(config, f)
}

func Save(config *Config, w io.Writer) error {
	endcoder := json.NewEncoder(w)
	if err := endcoder.Encode(config); err != nil {
		return err
	}
	return nil
}

// END OMIT

type Config struct {
	Name string `json:"name"`
}
