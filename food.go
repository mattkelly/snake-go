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
	var foodCanvas tl.Canvas
	if *isRealFood {
		runeSlice := []string{"🐁", "🐀", "🐇", "🐟", "🐠", "🐤", "🐥", "🐣", "🐸", "🐭", "🐰", "🥚"}
		victim := runeSlice[rand.Intn(len(runeSlice))]
		foodCanvas = tl.CanvasFromString(victim)
	} else {
		foodCanvas = tl.CanvasFromString("*")
	}

	f.Entity = tl.NewEntityFromCanvas(1, 1, foodCanvas)

	if !*isRealFood {
		f.Entity.SetCell(f.coord.x, f.coord.y, &tl.Cell{
			 		Fg: tl.ColorRed,
			 		Ch: '*',
			 	})
	}

	f.moveToRandomPosition()
	return f
}

// We don't really need to redraw the same item right now. 
// Will probably need this to draw new food items in a single Level though.
// Draw draws the Food as a default character.
//func (f *Food) Draw(screen *tl.Screen) {
	// if *isRealFood {
	// 	runeSlice := []rune{'🐁', '🐀', '🐇', '🐟', '🐠', '🐤', '🐥', '🐣', '🐸', '🐭', '🐰', '🥚'}
	// 	victim := runeSlice[rand.Intn(len(runeSlice))]
	// 	screen.RenderCell(f.coord.x, f.coord.y, &tl.Cell{
	// 		Fg: tl.ColorRed,
	// 		Ch: victim,
	// 	})
	// } else {
	// 	screen.RenderCell(f.coord.x, f.coord.y, &tl.Cell{
	// 		Fg: tl.ColorRed,
	// 		Ch: '*',
	// 	})
	// }
	//screen.RenderCell(f.Entity)
//}

// Position returns the x,y position of this Food.
func (f Food) Position() (int, int) {
	return f.coord.x, f.coord.y
}

// Size returns the size of this Food - always 1x1.
func (f Food) Size() (int, int) {
	// Emoji characters are wider and need a double hit box
	if *isRealFood{
		return 2, 1
	} else {
		return 1, 1
	}
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
