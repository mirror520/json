package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gitlab.com/mirror520/json/cors"
)

func main() {
	corsCfg, err := loadCORSRulesFromFile("./cors-rules.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(corsCfg)
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
