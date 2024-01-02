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

type Number struct {
	img  *ebiten.Image
	x, y float64
}

func loadNumber(path string) *Number {
	img := utils.ResizeImageFromReader(path, 1)
	return &Number{
		img: ebiten.NewImageFromImage(img),
	}
}

func (n *Number) draw(screen ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.x, n.y)
	screen.DrawImage(n.img, op)
}

func (n *Number) setXY(x, y float64) {
	n.x, n.y = x, y
}

var (
	numberMap = map[int]*Number{
		0: loadNumber("resource/war/number/number_0.png"),
		1: loadNumber("resource/war/number/number_1.png"),
		2: loadNumber("resource/war/number/number_2.png"),
		3: loadNumber("resource/war/number/number_3.png"),
		4: loadNumber("resource/war/number/number_4.png"),
		5: loadNumber("resource/war/number/number_5.png"),
		6: loadNumber("resource/war/number/number_6.png"),
		7: loadNumber("resource/war/number/number_7.png"),
		8: loadNumber("resource/war/number/number_8.png"),
		9: loadNumber("resource/war/number/number_9.png"),
	}
)
