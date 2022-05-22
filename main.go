package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"gitlab.com/mirror520/json/convert"
	"gitlab.com/mirror520/json/cors"
)

func main() {
	svc := convert.NewService()
	endpoint := convert.SnakeCaseEndpoint(svc)
	handler := convert.SnakeCaseHandler(endpoint)

	mux := http.NewServeMux()
	mux.Handle("/json/convert", handler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listen and Server in https://127.0.0.1:8080")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
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
