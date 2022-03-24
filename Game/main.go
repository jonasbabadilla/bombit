package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	game, _ := bombIt.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Rotate (Ebiten Demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
