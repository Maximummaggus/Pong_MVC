package main

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
		BallSpeedX:   ballSpeedX,
		BallSpeedY:   ballSpeedY,
	}
}

func (m *Model) Update() error {
	if !m.gameStarted {
		return nil
	}
	m.moveBall()
	m.checkWallCollision()
	m.checkPaddleCollision()
	m.resetBall()
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

func (m *Model) resetBall() {
	if m.BallX < -ballSize || m.BallX > screenWidth {
		*m = *NewModel()
		m.gameStarted = false
	}
}

func (m *Model) MovePaddle(up bool, left bool) {
	if left {
		if up {
			m.LeftPaddleY -= paddleSpeed
			if m.LeftPaddleY < 0 {
				m.LeftPaddleY = 0
			}
		} else {
			m.LeftPaddleY += paddleSpeed
			if m.LeftPaddleY > screenHeight-paddleHeight {
				m.LeftPaddleY = screenHeight - paddleHeight
			}
		}
	} else {
		if up {
			m.RightPaddleY -= paddleSpeed
			if m.RightPaddleY < 0 {
				m.RightPaddleY = 0
			}
		} else {
			m.RightPaddleY += paddleSpeed
			if m.RightPaddleY > screenHeight-paddleHeight {
				m.RightPaddleY = screenHeight - paddleHeight
			}
		}
	}
}
