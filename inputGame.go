package bombIt

import (

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// boom

type InputGame struct {
	pressed bool
}

func mouseInput() *InputGame {
	return &InputGame{}
}

// Update - display x, y coordinates
func (g *InputGame) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		var i InputGame
		i.pressed = true
	}
	return nil
}
