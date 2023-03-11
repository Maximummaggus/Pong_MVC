package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Konstanten für die Größe von Paddel und Ball sowie Bildschirmgröße
const (
	paddleWidth  = 16
	paddleHeight = 96
	ballSize     = 16
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	model := NewModel()
	view := NewView(model)
	controller := NewController(model, view)

	if err := ebiten.RunGame(controller); err != nil {
		panic(err)
	}
}
