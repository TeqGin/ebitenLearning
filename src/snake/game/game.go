package game

import (
	"ebitenLearning/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var _ ebiten.Game = &Game{}

type Direction int

type Game struct {
	s   *Snake
	f   *Fruit
	cfg *utils.Config
}

func NewGame() *Game {
	cfg := utils.LoadConfig("resource/snake/snake_config.json")
	ebiten.SetWindowSize(cfg.Width, cfg.Height)
	return &Game{
		s:   LoadSnake(),
		f:   LoadFruit("resource/snake/fruit/apple.png", 0.4, cfg),
		cfg: cfg,
	}
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	g.s.Draw(screen)
	g.f.Draw(screen)
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.cfg.Width, g.cfg.Height
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	g.s.Update(g)
	return nil
}
