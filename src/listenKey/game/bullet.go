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

func loadBullet(path string, cfg *config, a aircraft, speed float64, scalar float64) *bullet {
	var img = utils.ResizeImageFromReader(path, scalar)
	return &bullet{
		image: ebiten.NewImageFromImage(img),
		x:     a.getX() + float64(a.getImage().Bounds().Dx()-img.Bounds().Dx())/2,
		y:     a.getY() + float64(a.getImage().Bounds().Dy()-img.Bounds().Dy())/2,
		speed: speed,
	}
}

func (b *bullet) draw(screen *ebiten.Image, cfg *config, flipX float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	op.GeoM.Scale(flipX, 1)
	screen.DrawImage(b.image, op)
}

func (b *bullet) upadte() {
	b.y -= b.speed
}
