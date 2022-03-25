package bombIt

import (
	"bytes"
	"image"
	_ "image/png"

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jonasbabadilla/bombit/imageResources"
)

var g *Game

var (
	character *ebiten.Image
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var options *ebiten.DrawImageOptions

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
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	x, y := g.input.DirUp.x, g.input.DirUp.y
	options = &ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(-x), float64(-y))
	screen.DrawImage(character, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	img, _, _ := image.Decode(bytes.NewReader(imageResources.StartButton_png))

	character = ebiten.NewImageFromImage(img)

	return screenWidth, screenHeight
}

// mouse input
