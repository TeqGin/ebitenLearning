package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	i   *input
	cfg *config
	p   *plane
}

type input struct {
	msg string
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.Width, cfg.Hight)
	ebiten.SetWindowTitle(cfg.Title)
	return &Game{
		i:   &input{},
		cfg: cfg,
		p:   loadPlane("./resource/airplane/plane/plane1.png", cfg),
	}
}

func (i *input) update() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.msg = "go right"
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.msg = "go left"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		i.msg = "go space"
	}
}

// update the running data
func (g *Game) Update() error {
	g.p.update(g.cfg)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.p.Draw(screen, g.cfg)
	// ebitenutil.DebugPrint(screen, g.i.msg)
}

// logic size, use to zoom in/out the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.Width, g.cfg.Hight
}
