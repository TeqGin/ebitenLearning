package main

import (
	"ebitenLearning/src/listenKey/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
