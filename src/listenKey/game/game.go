package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	cfg *config
	p   *plane
	bg  *background
}

const (
	resourcePath = "resource"
)

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.Width, cfg.Hight)
	ebiten.SetWindowTitle(cfg.Title)
	return &Game{
		cfg: cfg,
		p:   loadPlane(resourcePath+"/airplane/plane/plane1.png", cfg),
		bg:  loadBackground(resourcePath + "/background/bg1.jpg"),
	}
}

// update the running data
func (g *Game) Update() error {
	g.bg.update()
	g.p.update(g.cfg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.bg.draw(screen, g.cfg)
	g.p.Draw(screen, g.cfg)
	// ebitenutil.DebugPrint(screen, g.i.msg)
}

// logic size, use to zoom in/out the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.Width, g.cfg.Hight
}
