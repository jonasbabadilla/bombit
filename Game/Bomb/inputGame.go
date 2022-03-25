package bombIt

import (

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// boom
type vec struct {
	x int
	y int
}

type InputGame struct {
	DirUp    vec
	DirDown  vec
	DirLeft  vec
	DirRight vec
}

func keyInput() *InputGame {

	return &InputGame{}
}

func (g *InputGame) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.DirUp = vec{0, 1}
	} else {
		g.DirUp = vec{0, 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.DirDown = vec{0, -1}
	} else {
		g.DirDown = vec{0, 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.DirLeft = vec{-1, 0}
	} else {
		g.DirLeft = vec{0, 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.DirRight = vec{1, 0}
	} else {
		g.DirRight = vec{0, 0}
	}

	return nil
}
