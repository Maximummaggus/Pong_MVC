package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	model := NewModel()
	view := NewView(model)
	controller := NewController(model, view)

	if err := ebiten.RunGame(controller); err != nil {
		panic(err)
	}
}
