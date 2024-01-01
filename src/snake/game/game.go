package game

import "github.com/hajimehoshi/ebiten/v2"

var _ ebiten.Game = &Game{}

type Direction int

type Game struct {
	s   *Snake
	f   *Fruit
	cfg *Config
}

func NewGame() *Game {
	cfg := LoadConfig()
	ebiten.SetWindowSize(cfg.Width, cfg.Hight)
	return &Game{
		s:   LoadSnake(),
		f:   LoadFruit("resource/fruit/strawberry.png", 0.025, cfg),
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
	return g.cfg.Width, g.cfg.Hight
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	g.s.Update(g)
	return nil
}
