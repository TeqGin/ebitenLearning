package main

import (
	"ebitenLearning/src/warGame/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
