package game

import (
	"ebitenLearning/src/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	boardImg          *ebiten.Image
	XImg              *ebiten.Image
	OImg              *ebiten.Image
	cfg               *utils.Config
	isX               bool
	lastClickInterval time.Time
	boardStatus       *BoardStaus
	winner            int
}

type BoardStaus struct {
	Pieces [][]int
}

func NewGame() *Game {
	cfg := utils.LoadConfig("resource/ooxx/config.json")
	ebiten.SetWindowSize(cfg.Width, cfg.Height)
	return &Game{
		cfg:      cfg,
		boardImg: utils.NewEbitenImangeFromFile("resource/ooxx/elements/board.png"),
		XImg:     utils.NewEbitenImangeFromFile("resource/ooxx/elements/x.png"),
		OImg:     utils.NewEbitenImangeFromFile("resource/ooxx/elements/o.png"),
		boardStatus: &BoardStaus{
			Pieces: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	var (
		img *ebiten.Image
		x   float64
		y   float64
	)
	width, height := g.XImg.Bounds().Dx()+4, g.XImg.Bounds().Dy()+6
	screen.DrawImage(g.boardImg, nil)
	for i := 0; i < len(g.boardStatus.Pieces); i++ {
		for j := 0; j < len(g.boardStatus.Pieces[i]); j++ {

			if g.boardStatus.Pieces[i][j] == 1 {
				img = g.OImg
			} else if g.boardStatus.Pieces[i][j] == 2 {
				img = g.XImg
			} else {
				continue
			}
			x = float64(width * i)
			y = float64(height * j)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			screen.DrawImage(img, op)
		}
	}
	if g.winner != 0 {
		w := ""
		if g.winner == 1 {
			w = "O"
		} else {
			w = "X"
		}
		ebitenutil.DebugPrintAt(screen, "last winner is "+w, 10, 620)
	}

}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.cfg.Width, g.cfg.Height
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	g.MouseClickListen()
	g.KeyBoardListen()

	if winer := g.CheckWiner(); winer != 0 {
		g.winner = winer
		g.Reset()
	}
	return nil
}

func (g *Game) MouseClickListen() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && time.Since(g.lastClickInterval).Milliseconds() > 500 {
		g.lastClickInterval = time.Now()
		mouseClickX, mouseClickY := ebiten.CursorPosition()
		i := mouseClickX / 200
		j := mouseClickY / 200
		if g.boardStatus.Pieces[i][j] != 0 {
			return
		}
		if g.isX {
			g.boardStatus.Pieces[i][j] = 2
		} else {
			g.boardStatus.Pieces[i][j] = 1
		}
		g.isX = !g.isX
	}
}

func (g *Game) KeyBoardListen() {
	var i, j int

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		i, j = 2, 0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		i, j = 1, 0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		i, j = 0, 0
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		i, j = 2, 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		i, j = 1, 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		i, j = 0, 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		i, j = 2, 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyX) {
		i, j = 1, 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		i, j = 0, 2
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.Reset()
		return
	} else {
		return
	}

	if g.boardStatus.Pieces[i][j] != 0 {
		return
	}

	if g.isX {
		g.boardStatus.Pieces[i][j] = 2
	} else {
		g.boardStatus.Pieces[i][j] = 1
	}
	g.isX = !g.isX
}

func (g *Game) Implements() ebiten.Game {
	return (*Game)(nil)
}

func (g *Game) Reset() {
	g.boardStatus.Pieces = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	g.isX = false
}

func (g *Game) CheckWiner() int {
	length := len(g.boardStatus.Pieces)
	for i := 0; i < length; i++ {
		n := g.boardStatus.Pieces[i][0]
		win := true
		if n == 0 {
			continue
		}
		for j := 1; j < length; j++ {
			if n != g.boardStatus.Pieces[i][j] {
				win = false
				break
			}
		}
		if win {
			return n
		}
	}

	n := g.boardStatus.Pieces[0][0]
	win := true
	for i := 1; i < length; i++ {
		if n != g.boardStatus.Pieces[i][i] {
			win = false
			break
		}
	}
	if win {
		return n
	}

	n = g.boardStatus.Pieces[0][length-1]
	win = true
	for i, j := 1, length-2; i < length; i, j = i+1, j-1 {
		if n != g.boardStatus.Pieces[i][j] {
			win = false
			break
		}
	}
	if win {
		return n
	}

	for i := 0; i < length; i++ {
		n := g.boardStatus.Pieces[0][i]
		win := true
		if n == 0 {
			continue
		}
		for j := 1; j < length; j++ {
			if n != g.boardStatus.Pieces[j][i] {
				win = false
				break
			}
		}
		if win {
			return n
		}
	}
	return 0
}
