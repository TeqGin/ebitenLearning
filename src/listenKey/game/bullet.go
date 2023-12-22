package game

import (
	"ebitenLearning/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type bullet struct {
	image *ebiten.Image
	x     float64
	y     float64
	speed float64
}

func loadBullet(path string, cfg *config, p *plane) *bullet {
	img := utils.ResizeImageFromReader(path, 0.2)
	// width := img.Bounds().Dx()
	// height := img.Bounds().Dy()

	return &bullet{
		image: ebiten.NewImageFromImage(img),
		x:     p.x + float64(p.image.Bounds().Dx())/2,
		y:     p.y,
		speed: 2,
	}
}

func (b *bullet) draw(screen *ebiten.Image, cfg *config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, op)
}

func (b *bullet) upadte() {
	b.y -= b.speed
}
