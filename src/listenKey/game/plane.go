package game

import (
	"ebitenLearning/src/utils"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type plane struct {
	image          *ebiten.Image
	x              float64
	y              float64
	bullets        []*bullet
	live           int
	speed          float64
	lastBulletTime time.Time
}

func loadPlane(path string, cfg *config) *plane {
	// img, _, err := ebitenutil.NewImageFromFile(path)
	img := utils.ResizeImageFromReader(path, 0.5)
	if img == nil {
		log.Fatal("resize image failed")
	}

	return &plane{
		image:   ebiten.NewImageFromImage(img),
		x:       float64(cfg.Width-img.Bounds().Dx()) / 2,
		y:       float64(cfg.Hight - img.Bounds().Dy()),
		live:    50,
		speed:   5,
		bullets: make([]*bullet, 0, 50),
	}
}

func (p *plane) Draw(screen *ebiten.Image, cfg *config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.image, op)
	for _, bullet := range p.bullets {
		bullet.draw(screen, cfg)
	}
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
	if ebiten.IsKeyPressed(ebiten.KeySpace) &&
		time.Since(p.lastBulletTime).Milliseconds() > cfg.BulletInterval {
		bullet := loadBullet(cfg, p)
		p.bullets = append(p.bullets, bullet)
		p.lastBulletTime = time.Now()
	}
	for _, bullet := range p.bullets {
		bullet.upadte()
	}
}
