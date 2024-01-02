package game

import (
	"bytes"
	"ebitenLearning/src/resource"
	"encoding/json"
	"image/color"
	"log"
)

type config struct {
	Width          int        `json:"width"`
	Hight          int        `json:"height"`
	Title          string     `json:"title"`
	BgColor        color.RGBA `json:"bgColor"`
	BulletInterval int64      `json:"bulletInterval"`
	EnemyInterval  int64      `json:"enemyInterval"`
}

func loadConfig() *config {
	var cfg config
	configBytes, err := resource.Asset("resource/war/war_config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(bytes.NewReader(configBytes)).Decode(&cfg); err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
