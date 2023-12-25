package game

import (
	"ebitenLearning/src/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	x, y  float64
	image *ebiten.Image
}

func loadIcon(path string, x, y, scalar float64) *Button {
	img := utils.ResizeImageFromReader(path, scalar)
	return &Button{
		x:     x,
		y:     y,
		image: ebiten.NewImageFromImage(img),
	}
}

func (b *Button) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, op)
}

func (b *Button) update(g *Game, status GameStatus) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && time.Since(g.lastClickInterval).Milliseconds() > 500 {
		g.lastClickInterval = time.Now()
		mouseClickX, mouseClickY := ebiten.CursorPosition()
		r1 := utils.Rectangle{
			Left:  utils.Point{X: b.x, Y: b.y},
			Right: utils.Point{X: b.x + float64(b.image.Bounds().Dx()), Y: b.y + float64(b.image.Bounds().Dy())},
		}
		r2 := utils.Rectangle{
			Left:  utils.Point{X: float64(mouseClickX), Y: float64(mouseClickY)},
			Right: utils.Point{X: float64(mouseClickX), Y: float64(mouseClickY)},
		}
		if utils.IsOverlappingPoint(r1, r2) {
			g.status = status
		}
	}
}
