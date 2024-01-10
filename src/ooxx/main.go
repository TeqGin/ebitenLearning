package main

import (
	"ebitenLearning/src/ooxx/game"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		fmt.Println(err)
	}
}
