package game

import (
	"bytes"
	"ebitenLearning/src/resource"
	"ebitenLearning/src/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Menu struct {
	bg             *background
	startGameImage *ebiten.Image
}

func loadMenu() *Menu {
	b, _ := resource.Asset("resource/icon/start.png")
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return &Menu{
		bg:             loadBackground(resourcePath+"/background/bg_sky_with_logo.jpg", 1.12),
		startGameImage: img,
	}
}

func (m *Menu) draw(screen *ebiten.Image, cfg *config) {
	m.bg.draw(screen, cfg)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(150, 300)
	screen.DrawImage(m.startGameImage, op)
}

func (m *Menu) update(g *Game) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseClickX, mouseClickY := ebiten.CursorPosition()
		r1 := utils.Rectangle{
			Left:  utils.Point{X: 150, Y: 300},
			Right: utils.Point{X: 150 + float64(m.startGameImage.Bounds().Dx()), Y: 300 + float64(m.startGameImage.Bounds().Dy())},
		}
		r2 := utils.Rectangle{
			Left:  utils.Point{X: float64(mouseClickX), Y: float64(mouseClickY)},
			Right: utils.Point{X: float64(mouseClickX), Y: float64(mouseClickY)},
		}
		if utils.IsOverlappingPoint(r1, r2) {
			g.status = RUNNING
		}
	}
}
