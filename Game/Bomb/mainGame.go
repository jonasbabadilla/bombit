package bombIt

import (
	_ "image/png"

	//"math"

	"fmt"

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
		input: keyInput(),
	}

	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()
	if g.input.DirUp {
		fmt.Println("Up")
	}
	if g.input.DirDown {
		fmt.Println("Down")
	}
	if g.input.DirLeft {
		fmt.Println("Left")
	}
	if g.input.DirRight {
		fmt.Println("Right")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// mouse input
