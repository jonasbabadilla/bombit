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

var (
	button *ebiten.Image
)

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
	w, h := button.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Move the image to the screen's center.
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	screen.DrawImage(button, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// mouse input
