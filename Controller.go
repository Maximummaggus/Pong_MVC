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
		c.pongModel.LeftPaddleUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.pongModel.LeftPaddleDown()
	}

	if ebiten.IsKeyPressed(ebiten.KeyI) {
		c.pongModel.RightPaddleUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyK) {
		c.pongModel.RightPaddleDown()
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
