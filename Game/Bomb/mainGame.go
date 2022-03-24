package bombIt

import (
	_ "image/png"

	"fmt"
	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var g *Game

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	input *InputGame
}

func NewGame() (*Game, error) {
	g = &Game{
		input: mouseInput(),
	}
	
	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()
	if g.input.pressed == true {
		fmt.Println("Pressed", g.input.mouseX, g.input.mouseY)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// mouse input
