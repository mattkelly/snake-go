package main

import tl "github.com/JoelOtter/termloop"
//import "fmt"

type Direction int

const (
	RIGHT Direction = iota
	LEFT
	UP
	DOWN
)

type Snake struct {
	*tl.Entity
	body []coord
	direction Direction
}

func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(1, 1, 1, 1)
	s.body = []coord {
		coord{3, 0},
		coord{4, 0},
		coord{5, 0}, // head
	}
	s.direction = RIGHT
	return s
}

func (snake *Snake) Head() *coord {
	return &snake.body[len(snake.body) - 1]
}

func (snake *Snake) UpdatePosition(x, y int) {
	// Update body
	for i := 0; i < len(snake.body) - 1; i++ {
		snake.body[i] = snake.body[i + 1]
	}

	// Update head
	snake.SetPosition(x, y) // position of Entity is just the head
	snake.Head().x, snake.Head().y = snake.Position()

}

func (snake *Snake) Draw(screen *tl.Screen) {
	// Draw snake
	for _, c := range snake.body {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Fg: tl.ColorBlack,
			Ch: 'o',
		})
	}
}

func (snake *Snake) Tick(event tl.Event) {
	// Find new direction
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			snake.direction = RIGHT
		case tl.KeyArrowLeft:
			snake.direction = LEFT
		case tl.KeyArrowUp:
			snake.direction = UP
		case tl.KeyArrowDown:
			snake.direction = DOWN
		}
	}

	// Move
	x, y := snake.Position()
	switch snake.direction {
	case RIGHT:
		snake.UpdatePosition(x + 1, y)
	case LEFT:
		snake.UpdatePosition(x - 1, y)
	case UP:
		snake.UpdatePosition(x, y - 1)
	case DOWN:
		snake.UpdatePosition(x, y + 1)
	}
}
