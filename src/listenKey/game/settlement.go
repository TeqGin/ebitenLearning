package game

import (
	"ebitenLearning/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Settlement struct {
	bg            *background
	restartButton *Button
	countWindows  *ebiten.Image
}

func loadSettlement(cfg *config) *Settlement {
	img := utils.ResizeImageFromReader("resource/background/settlement_failure.png", 1)
	return &Settlement{
		bg:            loadBackground("resource/background/bg_settlement.jpg", 0.8),
		restartButton: loadIcon("resource/icon/restart.png", 150, 460, 1),
		countWindows:  ebiten.NewImageFromImage(img),
	}
}

func (f *Settlement) draw(screen *ebiten.Image, cfg *config, score int) {
	f.bg.draw(screen, cfg)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100)
	screen.DrawImage(f.countWindows, op)

	drawNumber(screen, 250, 230, score/100)
	drawNumber(screen, 250, 300, score/20)
	drawNumber(screen, 250, 375, score)

	f.restartButton.draw(screen)
}

func (f *Settlement) update(g *Game, status GameStatus) {
	f.restartButton.update(g, status)
	if g.status == PREPARE {
		g.p.bullets = make(map[*bullet]struct{})
		g.enemies = make(map[*enemy]struct{})
		g.p.x = float64(g.cfg.Width-g.p.image.Bounds().Dx()) / 2
		g.p.y = float64(g.cfg.Hight - g.p.image.Bounds().Dy())
	}
}
