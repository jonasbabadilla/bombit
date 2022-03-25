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
	options = &ebiten.DrawImageOptions{}
	options.GeoM.Translate(screenWidth/2, screenHeight/2)
	return g, nil
}

func (g *Game) Update() error {
	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	var x, y float64
	x += float64(g.input.DirLeft+g.input.DirRight) 
	y += float64(g.input.DirUp+g.input.DirDown)
	
	options.GeoM.Translate(x, y)
	
	screen.DrawImage(character, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	img, _, _ := image.Decode(bytes.NewReader(imageResources.StartButton_png))

	character = ebiten.NewImageFromImage(img)
	

	return screenWidth, screenHeight
}

// mouse input
