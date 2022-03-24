package bombIt

import (
	_ "image/png"

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	input *InputGame
}

func NewGame() (*Game, error) {
	h := &Game{
		input: mouseInput(),
	}
	return h, nil
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// mouse input
