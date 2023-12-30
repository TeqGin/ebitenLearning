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
	X, Y          float64
}

func LoadFruit(path string, scalar float64) *Fruit {
	img := utils.ResizeImageFromReader(path, scalar)
	rand.NewSource(time.Now().UnixNano())
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	return &Fruit{
		Img:    ebiten.NewImageFromImage(img),
		Width:  w,
		Height: h,
		X:      float64(rand.Intn(500 - w)),
		Y:      float64(rand.Intn(500 - h)),
	}
}

func (f *Fruit) Generate() {
	f.X = float64(rand.Intn(500 - f.Width))
	f.Y = float64(rand.Intn(500 - f.Height))
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.X, f.Y)
	screen.DrawImage(f.Img, op)
}
