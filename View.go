package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strings"
)

type View struct {
	pongModel *Model
}

func NewView(model *Model) *View {
	return &View{
		pongModel: model,
	}
}

func drawText(screen *ebiten.Image, msg string, y int) {
	lines := strings.Split(msg, "\n")
	maxLineWidth := 0
	for _, line := range lines {
		if len(line) > maxLineWidth {
			maxLineWidth = len(line)
		}
	}
	textWidth := maxLineWidth * 16
	x := (screenWidth - textWidth) / 2

	for _, line := range lines {
		ebitenutil.DebugPrintAt(screen, line, x, y)
		y += 16
	}
}
func (v *View) Draw(screen *ebiten.Image) {
	if !v.pongModel.gameStarted {
		msg := "Control left Paddle -> W/S\nControl right Paddle -> I/K\n\nPress Enter to start"
		drawText(screen, msg, screenHeight/2)
	} else {
		ebitenutil.DrawRect(screen, 0, v.pongModel.LeftPaddleY, paddleWidth, paddleHeight, color.White)
		ebitenutil.DrawRect(screen, screenWidth-paddleWidth, v.pongModel.RightPaddleY, paddleWidth, paddleHeight, color.White)
		ebitenutil.DrawRect(screen, v.pongModel.BallX, v.pongModel.BallY, ballSize, ballSize, color.White)
	}
}

func (v *View) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}
