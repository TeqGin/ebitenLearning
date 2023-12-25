package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	bg          *background
	startButton *Button
}

func loadMenu() *Menu {
	return &Menu{
		bg:          loadBackground(resourcePath+"/background/bg_sky_with_logo.jpg", 1.12),
		startButton: loadIcon("resource/icon/start.png", 150, 300, 1),
	}
}

func (m *Menu) draw(screen *ebiten.Image, cfg *config) {
	m.bg.draw(screen, cfg)
	m.startButton.draw(screen)
}

func (m *Menu) update(g *Game) {
	m.startButton.update(g, RUNNING)
}
