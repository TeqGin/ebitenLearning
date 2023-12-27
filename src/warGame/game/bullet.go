package game

import (
	"ebitenLearning/src/utils"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type bullet struct {
	image     *ebiten.Image
	originX   float64
	originY   float64
	x         float64
	y         float64
	speed     float64
	isSpecial bool
	isRight   bool
}

func loadBullet(path string, cfg *config, a aircraft, speed float64, scalar float64, isSpecial, isRight bool) *bullet {
	var img = utils.ResizeImageFromReader(path, scalar)
	originX := a.getX() + float64(a.getImage().Bounds().Dx()-img.Bounds().Dx())/2
	originY := a.getY() + float64(a.getImage().Bounds().Dy()-img.Bounds().Dy())/2
	return &bullet{
		image:     ebiten.NewImageFromImage(img),
		originX:   originX,
		originY:   originY,
		x:         originX,
		y:         originY,
		speed:     speed,
		isSpecial: isSpecial,
		isRight:   isRight,
	}
}

func (b *bullet) draw(screen *ebiten.Image, cfg *config, flipX float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	if b.isSpecial {
		op.GeoM.Translate(-b.x, -b.y)
		sign := (b.x - b.originX) / math.Abs(b.x-b.originX)
		op.GeoM.Rotate(math.Atan2(sign, sign*0.03*(b.x-b.originX)))
		op.GeoM.Translate(b.x, b.y)
	}
	screen.DrawImage(b.image, op)
}

func (b *bullet) upadte() {
	if b.isSpecial {
		sign := float64(-1)
		if b.isRight {
			sign = 1
		}
		b.y -= b.speed
		b.x = sign*math.Sqrt(math.Abs((b.y-b.originY)/0.015)) + b.originX
	} else {
		b.y -= b.speed
	}
}
