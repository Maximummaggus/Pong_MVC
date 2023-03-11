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
	gameStarted  bool
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
	m.moveBall()
	m.checkWallCollision()
	m.checkPaddleCollision()
	m.resetBallIfOffScreen()
	return nil
}

func (m *Model) moveBall() {
	m.BallX += m.BallSpeedX
	m.BallY += m.BallSpeedY
}

func (m *Model) checkWallCollision() {
	if m.BallY <= 0 || m.BallY >= screenHeight-ballSize {
		m.BallSpeedY = -m.BallSpeedY
	}
}

func (m *Model) checkPaddleCollision() {
	ballCenterX := m.BallX + ballSize/2
	ballCenterY := m.BallY + ballSize/2
	if ballCenterX <= paddleWidth && ballCenterY >= m.LeftPaddleY && ballCenterY <= m.LeftPaddleY+paddleHeight {
		m.BallSpeedX = -m.BallSpeedX
	}
	if ballCenterX >= screenWidth-paddleWidth-ballSize && ballCenterY >= m.RightPaddleY && ballCenterY <= m.RightPaddleY+paddleHeight {
		m.BallSpeedX = -m.BallSpeedX
	}
}

func (m *Model) resetBallIfOffScreen() {
	if m.BallX < -ballSize || m.BallX > screenWidth {
		m.BallX = screenWidth / 2
		m.BallY = screenHeight / 2
		rand.Seed(time.Now().UnixNano())
		m.BallSpeedX = 7
		m.BallSpeedY = 7
		m.gameStarted = false
	}
}

func (m *Model) MovePaddle(up bool, left bool) {
	if left {
		if up {
			m.LeftPaddleY -= 5
			if m.LeftPaddleY < 0 {
				m.LeftPaddleY = 0
			}
		} else {
			m.LeftPaddleY += 5
			if m.LeftPaddleY > screenHeight-paddleHeight {
				m.LeftPaddleY = screenHeight - paddleHeight
			}
		}
	} else {
		if up {
			m.RightPaddleY -= 5
			if m.RightPaddleY < 0 {
				m.RightPaddleY = 0
			}
		} else {
			m.RightPaddleY += 5
			if m.RightPaddleY > screenHeight-paddleHeight {
				m.RightPaddleY = screenHeight - paddleHeight
			}
		}
	}
}
