package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type plane struct {
	image *ebiten.Image
	x     float64
	y     float64
	live  int
	speed float64
}

func loadPlane(path string, cfg *config) *plane {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return &plane{
		image: img,
		x:     float64(cfg.Width-img.Bounds().Dx()) / 2,
		y:     float64(cfg.Hight - img.Bounds().Dy()),
		live:  50,
		speed: 5,
	}
}

func (p *plane) Draw(screen *ebiten.Image, cfg *config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, op)
}

func (p *plane) update(cfg *config) {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += p.speed
		boundary := float64(cfg.Width - p.image.Bounds().Dx())
		if p.x >= boundary {
			p.x = boundary
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= p.speed
		if p.x <= 0 {
			p.x = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y -= p.speed
		if p.y <= 0 {
			p.y = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y += p.speed
		boundary := float64(cfg.Hight - p.image.Bounds().Dy())
		if p.y > boundary {
			p.y = boundary
		}
	}
}
