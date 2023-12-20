package game

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type config struct {
	Width   int        `json:"width"`
	Hight   int        `json:"height"`
	Title   string     `json:"title"`
	BgColor color.RGBA `json:"bgColor"`
}

func loadConfig() *config {
	f, err := os.Open("./resource/config.json")
	if err != nil {
		log.Fatalf("open config failed: %v", err)
	}

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("decode config failed: %v", err)
	}

	return &cfg
}
