package game

import (
	"ebitenLearning/src/utils"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

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
	Body     []utils.Point
	Dir      Direction
	timer    int
	moveTime int
}

func LoadSnake() *Snake {
	return &Snake{
		Body:     []utils.Point{{X: 5, Y: 5}},
		Dir:      DIR_NONE,
		moveTime: 3,
	}
}

func (s *Snake) Update(g *Game) {
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
		if s.IsCollisionWithWall(g.cfg) || s.IsCollisionWithSelf() {
			return
		}

		if s.IsCollisionWithFruit(g) {
			s.Body = append(s.Body, s.Body[len(s.Body)-1])
		}

		for i := len(s.Body) - 1; i > 0; i-- {
			s.Body[i].X = s.Body[i-1].X
			s.Body[i].Y = s.Body[i-1].Y
		}

		switch s.Dir {
		case UP:
			s.Body[0].Y--
		case DOWN:
			s.Body[0].Y++
		case LEFT:
			s.Body[0].X--
		case RIGHT:
			s.Body[0].X++
		}
	}

	s.timer++
}

func (s *Snake) needMove() bool {
	return s.timer%s.moveTime == 0
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, v := range s.Body {
		vector.DrawFilledRect(screen, float32(v.X*gridSize), float32(v.Y*gridSize), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)
	}
}

func (s *Snake) Reset() {
	s.Body = []utils.Point{{X: 5, Y: 5}}
	s.timer = 0
	s.moveTime = 3
	s.Dir = DIR_NONE
}

func (s *Snake) IsCollisionWithWall(cfg *Config) bool {
	head := s.Body[0]
	if head.X <= 0 || head.X >= float64(cfg.Width/gridSize) ||
		head.Y <= 0 || head.Y >= float64(cfg.Hight/gridSize) {
		return true
	}
	return false
}

func (s *Snake) IsCollisionWithFruit(g *Game) bool {
	head := s.Body[0]
	if head.X == float64(g.f.X) && head.Y == float64(g.f.Y) {
		g.f.Generate(g.cfg)
		return true
	}
	return false
}

func (s *Snake) IsCollisionWithSelf() bool {
	head := s.Body[0]
	for i := 1; i < len(s.Body); i++ {
		if head.X == s.Body[i].X &&
			head.Y == s.Body[i].Y {
			return true
		}
	}
	return false
}
