package game

import (
	"ebitenLearning/src/utils"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fruit struct {
	Img           *ebiten.Image
	Width, Height int
	X, Y          int
}

func LoadFruit(path string, scalar float64, cfg *Config) *Fruit {
	img := utils.ResizeImageFromReader(path, scalar)
	rand.NewSource(time.Now().UnixNano())
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	return &Fruit{
		Img:    ebiten.NewImageFromImage(img),
		Width:  w,
		Height: h,
		X:      rand.Intn(cfg.Width / gridSize),
		Y:      rand.Intn(cfg.Hight / gridSize),
	}
}

func (f *Fruit) Generate(cfg *Config) {
	f.X = rand.Intn(cfg.Width / gridSize)
	f.Y = rand.Intn(cfg.Hight / gridSize)
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	if f.X < 0 || f.Y < 0 {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(f.X*gridSize), float64(f.Y*gridSize))
	screen.DrawImage(f.Img, op)
}

func (f *Fruit) Remove() {
	f.X = -1
	f.Y = -1
}
