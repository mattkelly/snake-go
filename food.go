package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

type Food struct {
	*tl.Entity
	coord Coord
}

func NewFood() *Food {
	f := new(Food)
	f.Entity = tl.NewEntity(1, 1, 1, 1)
	f.moveToRandomPosition()
	return f
}

func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.coord.x, food.coord.y, &tl.Cell{
		Fg: tl.ColorGreen,
		Ch: '*',
	})
}

func (food Food) Position() (int, int) {
	return food.coord.x, food.coord.y
}

func (food Food) Size() (int, int) {
	return 1, 1
}

func (food *Food) moveToRandomPosition() {
	// TODO actual range
	newX := randInRange(1, 50)
	newY := randInRange(1, 50)
	food.coord.x, food.coord.y = newX, newY
	food.SetPosition(newX, newY)
}

func (food *Food) handleSnakeCollision() {
	food.moveToRandomPosition()
	score++
}

func (food *Food) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Snake:
		// It better be a snake that we're colliding with...
		food.handleSnakeCollision()
	}
}

func randInRange(min, max int) int {
	return rand.Intn(max - min) + min
}
