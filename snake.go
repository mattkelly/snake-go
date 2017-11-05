package main

import tl "github.com/JoelOtter/termloop"

type Direction int

const (
	RIGHT Direction = iota
	LEFT
	UP
	DOWN
)

type Snake struct {
	*tl.Entity
	body []Coord
	bodyLen int
	direction Direction
}

func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(5, 5, 1, 1)
	s.body = []Coord {
		Coord{3, 5},
		Coord{4, 5},
		Coord{5, 5}, // head
	}
	// Need to track length explicitly for the case
	// where we're actively growing
	s.bodyLen = len(s.body)
	s.direction = RIGHT
	return s
}

func (s *Snake) Head() *Coord {
	return &s.body[len(s.body) - 1]
}

func (s *Snake) grow(amount int) {
	s.bodyLen += amount
}

func (s *Snake) isGrowing() bool {
	return s.bodyLen > len(s.body)
}

func (s *Snake) isCollidingWithSelf() bool {
	for i := 0; i < len(s.body) - 1; i++ {
		if *s.Head() == s.body[i] {
			return true
		}
	}
	return false
}

func (s *Snake) isCollidingWithBorder() bool {
	return border.Contains(*s.Head())
}

// Draw() is called every frame, whereas Tick() is
// only called on events.
func (s *Snake) Draw(screen *tl.Screen) {
	// Update position based on direction
	newHead := *s.Head()
	switch s.direction {
	case RIGHT:
		newHead.x += 1
	case LEFT:
		newHead.x -= 1
	case UP:
		newHead.y -= 1
	case DOWN:
		newHead.y += 1
	}

	if s.isGrowing() {
		// We must be growing
		s.body = append(s.body, newHead)
	} else {
		s.body = append(s.body[1:], newHead)
	}

	s.SetPosition(newHead.x, newHead.y)

	if s.isCollidingWithSelf() || s.isCollidingWithBorder() {
		EndGame()
	}

	// Draw snake
	for _, c := range s.body {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: 'o',
		})
	}
}

func (s *Snake) Tick(event tl.Event) {
	// Find new direction - but you can't go
	// back from where you came.
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if s.direction != LEFT {
				s.direction = RIGHT
			}
		case tl.KeyArrowLeft:
			if s.direction != RIGHT {
				s.direction = LEFT
			}
		case tl.KeyArrowUp:
			if s.direction != DOWN {
				s.direction = UP
			}
		case tl.KeyArrowDown:
			if s.direction != UP {
				s.direction = DOWN
			}
		}
	}
}

func (s *Snake) handleFoodCollision() {
	s.grow(1)
}

func (s *Snake) handleBorderCollision() {
	EndGame()
}

func (s *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		s.handleFoodCollision()
	case *Border:
		s.handleBorderCollision()
	}
}
