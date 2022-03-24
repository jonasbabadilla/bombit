package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	//"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jonasbabadilla/bombit/imageResources"
)

type MouseButton = driver.MouseButton

const (
	MouseButtonLeft   MouseButton = driver.MouseButtonLeft
	MouseButtonRight  MouseButton = driver.MouseButtonRight
	MouseButtonMiddle MouseButton = driver.MouseButtonMiddle
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
}

var (
	button *ebiten.Image
)

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
func IsMouseButtonPressed(mouseButton MouseButton) bool {
	return uiDriver().Input().IsMouseButtonPressed(mouseButton)
}

func CursorPosition() (x, y int) {
	return uiDriver().Input().CursorPosition()
}


func main() {

	img, _, err := image.Decode(bytes.NewReader(imageResources.StartButton_png))
	if err != nil {
		log.Fatal(err)
	}
	button = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rotate (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
