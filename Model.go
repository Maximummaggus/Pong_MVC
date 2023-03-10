package main

import (
	"math/rand"
	"time"
)

type Model struct {
	LeftPaddleY  float64
	RightPaddleY float64
	BallX        float64
	BallY        float64
	BallSpeedX   float64
	BallSpeedY   float64
	gameStarted  bool // Add this line
}

func NewModel() *Model {
	return &Model{
		LeftPaddleY:  screenHeight / 2,
		RightPaddleY: screenHeight / 2,
		BallX:        screenWidth / 2,
		BallY:        screenHeight / 2,
		BallSpeedX:   5,
		BallSpeedY:   5,
	}
}

func (m *Model) Update() error {
	if !m.gameStarted {
		return nil
	}

	// Ball movement
	m.BallX += m.BallSpeedX
	m.BallY += m.BallSpeedY

	// Check for wall collision
	if m.BallY <= 0 || m.BallY >= screenHeight-ballSize {
		m.BallSpeedY = -m.BallSpeedY
	}

	// Left paddle collision
	if m.BallX <= paddleWidth && m.BallY >= m.LeftPaddleY && m.BallY <= m.LeftPaddleY+paddleHeight {
		m.BallSpeedX = -m.BallSpeedX
	}

	// Right paddle collision
	if m.BallX >= screenWidth-paddleWidth-ballSize && m.BallY >= m.RightPaddleY && m.BallY <= m.RightPaddleY+paddleHeight {
		m.BallSpeedX = -m.BallSpeedX
	}

	// Reset ball position and pause game if it goes off-screen
	if m.BallX < -ballSize ||
		m.BallX > screenWidth {

		// Reset ball position and velocity
		m.BallX = screenWidth / 2
		m.BallY = screenHeight / 2

		rand.Seed(time.Now().UnixNano())

		m.BallSpeedX = 7

		m.BallSpeedY = 7

		// Set gameStarted to false to wait for Enter key press
		m.gameStarted = false

		return nil
	}

	return nil
}

func (m *Model) LeftPaddleUp() {
	m.LeftPaddleY -= 5
	if m.LeftPaddleY < 0 {
		m.LeftPaddleY = 0
	}
}

func (m *Model) LeftPaddleDown() {
	m.LeftPaddleY += 5
	if m.LeftPaddleY > screenHeight-paddleHeight {
		m.LeftPaddleY = screenHeight - paddleHeight
	}
}

func (m *Model) RightPaddleUp() {
	m.RightPaddleY -= 5
	if m.RightPaddleY < 0 {
		m.RightPaddleY = 0
	}
}

func (m *Model) RightPaddleDown() {
	m.RightPaddleY += 5
	if m.RightPaddleY > screenHeight-paddleHeight {
		m.RightPaddleY = screenHeight - paddleHeight
	}
}
