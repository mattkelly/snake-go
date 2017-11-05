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

func (snake *Snake) Head() *Coord {
	return &snake.body[len(snake.body) - 1]
}

func (snake *Snake) Grow(amount int) {
	snake.bodyLen += amount
}

func (snake *Snake) IsGrowing() bool {
	return snake.bodyLen > len(snake.body)
}

// Draw() is called every frame, whereas Tick() is
// only called on events.
func (snake *Snake) Draw(screen *tl.Screen) {
	// Update position based on direction
	newHead := *snake.Head()
	switch snake.direction {
	case RIGHT:
		newHead.x += 1
	case LEFT:
		newHead.x -= 1
	case UP:
		newHead.y -= 1
	case DOWN:
		newHead.y += 1
	}

	if snake.IsGrowing() {
		// We must be growing
		snake.body = append(snake.body, newHead)
	} else {
		snake.body = append(snake.body[1:], newHead)
	}

	snake.SetPosition(newHead.x, newHead.y)

	// Draw snake
	for _, c := range snake.body {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: 'o',
		})
	}
}

func (snake *Snake) Tick(event tl.Event) {
	// Find new direction - but you can't go
	// back from where you came.
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.direction != LEFT {
				snake.direction = RIGHT
			}
		case tl.KeyArrowLeft:
			if snake.direction != RIGHT {
				snake.direction = LEFT
			}
		case tl.KeyArrowUp:
			if snake.direction != DOWN {
				snake.direction = UP
			}
		case tl.KeyArrowDown:
			if snake.direction != UP {
				snake.direction = DOWN
			}
		}
	}
}

func (snake *Snake) handleFoodCollision() {
	snake.Grow(1)
}

func (snake *Snake) handleBorderCollision() {
	// dead
}

func (snake *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		snake.handleFoodCollision()
	case *Border:
		snake.handleBorderCollision()
	}
}
