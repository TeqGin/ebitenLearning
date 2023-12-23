package game

import (
	"ebitenLearning/src/utils"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type enemy struct {
	image   *ebiten.Image
	x       float64
	y       float64
	speed   float64
	bullets []*bullet
}

var enemyImg = utils.ResizeImageFromReader("resource/airplane/enemy/enemy1.png", 0.15)

func loadEnemy(cfg *config) *enemy {
	rand.NewSource(time.Now().UnixNano())

	return &enemy{
		image:   ebiten.NewImageFromImage(enemyImg),
		x:       float64(rand.Intn(cfg.Width - enemyImg.Bounds().Dx())),
		y:       -float64(enemyImg.Bounds().Dy()),
		speed:   3,
		bullets: make([]*bullet, 0, 10),
	}
}

func (e *enemy) draw(screen *ebiten.Image, cfg *config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(e.image, op)
}

func (e *enemy) update() {
	e.y += e.speed
}