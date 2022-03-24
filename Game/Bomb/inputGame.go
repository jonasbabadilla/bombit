package bombIt

import (

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// boom

type InputGame struct {
	pressed bool
	mouseX int
	mouseY int
}

func mouseInput() *InputGame {
	return &InputGame{}
}


func (g *InputGame) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.pressed = true
		g.mouseX, g.mouseY = ebiten.CursorPosition()
	} else {
		g.pressed = false
	}
	return nil
}
