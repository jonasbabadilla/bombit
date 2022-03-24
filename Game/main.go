package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	bombIt "github.com/jonasbabadilla/bombit/Game/Bomb"
)

func main() {
	//testing
	game, _ := bombIt.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Rotate (Ebiten Demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
