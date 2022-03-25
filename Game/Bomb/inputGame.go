package bombIt

import (

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// boom

type InputGame struct {
	DirUp, DirDown, DirLeft, DirRight int
}

func keyInput() *InputGame {

	return &InputGame{}
}

func (g *InputGame) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.DirUp = -1
	} else {
		g.DirUp = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.DirDown = 1
	} else {
		g.DirDown = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.DirLeft = -1
	} else {
		g.DirLeft = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.DirRight = 1
	} else {
		g.DirRight = 0
	}
	return nil
}
