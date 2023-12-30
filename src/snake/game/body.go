package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point struct {
	x, y int
}

const (
	DIR_NONE Direction = iota
	UP
	DOWN
	LEFT
	RIGHT
)

const (
	gridSize = 10
)

type Snake struct {
	Body         []Point
	Dir          Direction
	lastGrowTime time.Time
	timer        int
	moveTime     int
}

func LoadSnake() *Snake {
	return &Snake{
		Body:     []Point{{5, 5}},
		Dir:      DIR_NONE,
		moveTime: 3,
	}
}

func (s *Snake) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if s.Dir == DOWN {
			return
		}
		s.Dir = UP
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if s.Dir == UP {
			return
		}
		s.Dir = DOWN
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if s.Dir == RIGHT {
			return
		}
		s.Dir = LEFT
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if s.Dir == LEFT {
			return
		}
		s.Dir = RIGHT
	}

	if s.needMove() {
		if s.Dir != DIR_NONE && time.Since(s.lastGrowTime).Milliseconds() > 1000 {
			s.Body = append(s.Body, s.Body[len(s.Body)-1])
			s.lastGrowTime = time.Now()
		}

		for i := len(s.Body) - 1; i > 0; i-- {
			s.Body[i].x = s.Body[i-1].x
			s.Body[i].y = s.Body[i-1].y
		}

		switch s.Dir {
		case UP:
			s.Body[0].y--
		case DOWN:
			s.Body[0].y++
		case LEFT:
			s.Body[0].x--
		case RIGHT:
			s.Body[0].x++
		}
	}

	s.timer++
}

func (s *Snake) needMove() bool {
	return s.timer%s.moveTime == 0
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, v := range s.Body {
		vector.DrawFilledRect(screen, float32(v.x*gridSize), float32(v.y*gridSize), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)
	}
}
