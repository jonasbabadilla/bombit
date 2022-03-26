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

		g.DirUp = -2
	} else {
		g.DirUp = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.DirDown = 2
	} else {
		g.DirDown = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.DirLeft = -2
	} else {
		g.DirLeft = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.DirRight = 2
	} else {
		g.DirRight = 0
	}
	return nil
}
