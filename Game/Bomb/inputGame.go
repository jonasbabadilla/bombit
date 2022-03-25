package bombIt

import (

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// boom

type InputGame struct {
	DirUp    bool
	DirDown  bool
	DirLeft  bool
	DirRight bool
}

func keyInput() *InputGame {

	return &InputGame{}
}

func (g *InputGame) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.DirUp = true
	} else {
		g.DirUp = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.DirDown = true
	} else {
		g.DirDown = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.DirLeft = true
	} else {
		g.DirLeft = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.DirRight = true
	} else {
		g.DirRight = false
	}

	return nil
}
