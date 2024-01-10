package utils

import (
	"bytes"
	"ebitenLearning/src/resource"
	"encoding/json"
	"log"
)

type Config struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Title  string `json:"title"`
}

func LoadConfig(path string) *Config {
	var cfg Config
	configBytes, err := resource.Asset(path)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(bytes.NewReader(configBytes)).Decode(&cfg); err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
