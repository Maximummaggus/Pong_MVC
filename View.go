package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type View struct {
	pongModel *Model
}

func NewView(model *Model) *View {
	return &View{
		pongModel: model,
	}
}

func (v *View) Draw(screen *ebiten.Image) {
	if !v.pongModel.gameStarted {
		msg := "Press Enter to start"
		textWidth := len(msg) * 16
		x := (screenWidth - textWidth) / 2
		y := screenHeight / 2

		ebitenutil.DebugPrintAt(screen, msg, x, y)
	} else {
		ebitenutil.DrawRect(screen, 0, v.pongModel.LeftPaddleY, paddleWidth, paddleHeight, color.White)
		ebitenutil.DrawRect(screen, screenWidth-paddleWidth, v.pongModel.RightPaddleY, paddleWidth, paddleHeight, color.White)
		ebitenutil.DrawRect(screen, v.pongModel.BallX, v.pongModel.BallY, ballSize, ballSize, color.White)
	}
}

func (v *View) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}
