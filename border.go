package main

import tl "github.com/JoelOtter/termloop"

// Border is the edge of the playing area. If the Snake collides with it,
// it dies.
type Border struct {
	*tl.Entity
	width, height int
	coords        map[Coord]int
}

// NewBorder creates a Border with the given dimensions.
func NewBorder(width, height int) *Border {
	b := new(Border)
	b.Entity = tl.NewEntity(1, 1, 1, 1)
	b.width, b.height = width, height

	b.coords = make(map[Coord]int)

	// Top and bottom
	for x := 0; x < width; x++ {
		b.coords[Coord{x, 0}] = 1
		b.coords[Coord{x, b.height}] = 1
	}

	// Left and right
	for y := 0; y < height+1; y++ {
		b.coords[Coord{0, y}] = 1
		b.coords[Coord{b.width, y}] = 1
	}

	return b
}

// Contains returns true if a Coord is part of the border, else false.
// Used for collision detection.
func (b *Border) Contains(coord Coord) bool {
	_, exists := b.coords[coord]
	return exists
}

// Draw draws the border on the screen. A default color is used.
func (b *Border) Draw(screen *tl.Screen) {
	if b == nil {
		return
	}

	for c := range b.coords {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}
}
