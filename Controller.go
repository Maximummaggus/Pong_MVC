package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Controller struct {
	pongModel   *Model
	pongView    *View
	gameStarted bool
}

func NewController(model *Model, view *View) *Controller {
	return &Controller{
		pongModel:   model,
		pongView:    view,
		gameStarted: false,
	}
}

func (c *Controller) Update() error {
	if !c.pongModel.gameStarted {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			c.pongModel.gameStarted = true
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.pongModel.MovePaddle(true, true)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.pongModel.MovePaddle(false, true)
	}

	if ebiten.IsKeyPressed(ebiten.KeyI) {
		c.pongModel.MovePaddle(true, false)
	}

	if ebiten.IsKeyPressed(ebiten.KeyK) {
		c.pongModel.MovePaddle(false, false)
	}

	if err := c.pongModel.Update(); err != nil {
		return err
	}

	return nil
}

func (c *Controller) Draw(screen *ebiten.Image) {
	c.pongView.Draw(screen)
}

func (c *Controller) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}
