package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

// Food handles its own collisions with a Snake and places itself randomly
// on the screen. Since there's no top-level controller, it also updates the
// score.
type Food struct {
	*tl.Entity
	coord Coord
}

// NewFood creates a new Food at a random position.
func NewFood() *Food {
	f := new(Food)
	f.Entity = tl.NewEntity(1, 1, 1, 1)
	f.moveToRandomPosition()
	return f
}

// Draw draws the Food as a default character.
func (f *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(f.coord.x, f.coord.y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '*',
	})
}

// Position returns the x,y position of this Food.
func (f Food) Position() (int, int) {
	return f.coord.x, f.coord.y
}

// Size returns the size of this Food - always 1x1.
func (f Food) Size() (int, int) {
	return 1, 1
}

// Collide handles collisions with the Snake. It updates the score and places
// the Food randomly on the screen again.
func (f *Food) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Snake:
		// It better be a snake that we're colliding with...
		f.handleSnakeCollision()
	}
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

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
