package game

import (
	"bytes"
	"ebitenLearning/src/resource"
	"encoding/json"
	"log"
)

type Config struct {
	Width int    `json:"width"`
	Hight int    `json:"height"`
	Title string `json:"title"`
}

func LoadConfig() *Config {
	var cfg Config
	configBytes, err := resource.Asset("resource/snake/snake_config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(bytes.NewReader(configBytes)).Decode(&cfg); err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
