package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	i   *input
	cfg *config
}

type input struct {
	msg string
}

func NewGame() *Game {
	return &Game{
		i:   &input{},
		cfg: loadConfig(),
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
	g.i.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	ebitenutil.DebugPrint(screen, g.i.msg)
}

// logic size, use to zoom in/out the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640 / 2, 480 / 2
}
