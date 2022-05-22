package main

import (
	"encoding/json"
	"os"

	"gitlab.com/mirror520/json/cors"
)

func main() {
}

func loadCORSRulesFromFile(filename string) (*cors.Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var cfg cors.Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
