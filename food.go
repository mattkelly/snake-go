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

func (f *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(f.coord.x, f.coord.y, &tl.Cell{
		Fg: tl.ColorGreen,
		Ch: '*',
	})
}

func (f Food) Position() (int, int) {
	return f.coord.x, f.coord.y
}

func (f Food) Size() (int, int) {
	return 1, 1
}

func (f *Food) moveToRandomPosition() {
	newX := randInRange(1, border.width-1)
	newY := randInRange(1, border.height-1)
	f.coord.x, f.coord.y = newX, newY
	f.SetPosition(newX, newY)
}

func (f *Food) handleSnakeCollision() {
	f.moveToRandomPosition()
	IncreaseScore(5)
}

func (f *Food) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Snake:
		// It better be a snake that we're colliding with...
		f.handleSnakeCollision()
	}
}

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
