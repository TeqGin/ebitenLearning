package game

import "github.com/hajimehoshi/ebiten/v2"

var _ ebiten.Game = &Game{}

type Direction int

type Game struct {
	s *Snake
}

func NewGame() *Game {
	ebiten.SetWindowSize(500, 500)
	return &Game{
		s: LoadSnake(),
	}
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	g.s.Draw(screen)
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 500, 500
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	g.s.Update()
	return nil
}
