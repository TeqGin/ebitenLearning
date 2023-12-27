package game

import (
	"bytes"
	"ebitenLearning/src/resource"
	"ebitenLearning/src/utils"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type enemy struct {
	image          *ebiten.Image
	x              float64
	y              float64
	speed          float64
	lastLoadBullet time.Time
	bullets        []*bullet
	bloomPlayer    *audio.Player
}

var enemyImg = utils.ResizeImageFromReader("resource/airplane/enemy/enemy1.png", 0.15)
var audioContext *audio.Context

func loadEnemy(cfg *config) *enemy {
	rand.NewSource(time.Now().UnixNano())

	if audioContext == nil {
		sampleRate := 48000
		audioContext = audio.NewContext(sampleRate)
	}
	// Decode wav-formatted data and retrieve decoded PCM stream.
	b, _ := resource.Asset("resource/music/bloom.wav")
	d, _ := wav.DecodeWithoutResampling(bytes.NewReader(b))

	// Create an audio.Player that has one stream.
	audioPlayer, _ := audioContext.NewPlayer(d)
	return &enemy{
		image:       ebiten.NewImageFromImage(enemyImg),
		x:           float64(rand.Intn(cfg.Width - enemyImg.Bounds().Dx())),
		y:           -float64(enemyImg.Bounds().Dy()),
		speed:       3,
		bullets:     make([]*bullet, 0, 10),
		bloomPlayer: audioPlayer,
	}
}

func (e *enemy) draw(screen *ebiten.Image, cfg *config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(e.image, op)
	for _, bullet := range e.bullets {
		bullet.draw(screen, cfg, 1)
	}
}

func (e *enemy) update(cfg *config) {
	e.y += e.speed
	if time.Since(e.lastLoadBullet).Milliseconds() > 500 {
		bullet := loadBullet("resource/airplane/bullet/enemy_bullet2.png", cfg, e, -4, 1, false, false)
		e.bullets = append(e.bullets, bullet)
		e.lastLoadBullet = time.Now()
	}
	for _, bullet := range e.bullets {
		bullet.upadte()
	}
}

func (e *enemy) getX() float64 {
	return e.x
}

func (e *enemy) getY() float64 {
	return e.y
}

func (e *enemy) getImage() *ebiten.Image {
	return e.image
}
