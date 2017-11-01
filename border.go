package main

import tl "github.com/JoelOtter/termloop"

type Border struct {
	*tl.Entity
	width int
	height int
}

func NewBorder() *Border {
	b := new(Border)
	b.Entity = tl.NewEntity(1, 1, 1, 1)
	return b
}

func (border *Border) Draw(screen *tl.Screen) {
	// Draw top and bottom
	width, height := screen.Size()
	for x := 0; x < width; x++ {
		screen.RenderCell(x, 0, &tl.Cell{
			Bg: tl.ColorBlue,
		})
		screen.RenderCell(x, height - 1, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}

	// Draw left and right sides
	for y := 0; y < height; y++ {
		screen.RenderCell(0, y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
		screen.RenderCell(width - 1, y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}
}
