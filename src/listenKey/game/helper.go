package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func drawNumber(screen *ebiten.Image, x, y float64, num int) {
	numbers := make([]int, 0, 10)
	for num > 0 {
		numbers = append(numbers, num%10)
		num /= 10
	}
	if len(numbers) == 0 {
		numbers = append(numbers, 0)
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		number := numberMap[numbers[i]]
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(number.img, op)
		x += float64(number.img.Bounds().Dx()) + 5
	}
}
