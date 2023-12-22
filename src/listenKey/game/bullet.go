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

var img = utils.ResizeImageFromReader("resource/airplane/bullet/bullet1.png", 0.2)

func loadBullet(cfg *config, p *plane) *bullet {
	return &bullet{
		image: ebiten.NewImageFromImage(img),
		x:     p.x + float64(p.image.Bounds().Dx())/2,
		y:     p.y,
		speed: 6,
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
