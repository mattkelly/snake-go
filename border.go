package main

import tl "github.com/JoelOtter/termloop"

type Border struct {
	*tl.Entity
	width int
	height int
}

// TODO Draw() knows about screen anyway so use screen directly
// for height/width?
func NewBorder(width, height int) *Border {
	b := new(Border)
	b.Entity = tl.NewEntity(1, 1, 1, 1)
	b.width, b.height = width, height
	return b
}

func (border *Border) Draw(screen *tl.Screen) {
	// Draw top and bottom
	for x := 0; x < border.width; x++ {
		screen.RenderCell(x, 0, &tl.Cell{
			Bg: tl.ColorBlue,
		})
		screen.RenderCell(x, border.height - 1, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}

	// Draw left and right sides
	for y := 0; y < border.height; y++ {
		screen.RenderCell(0, y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
		screen.RenderCell(border.width - 1, y, &tl.Cell{
			Bg: tl.ColorBlue,
		})
	}
}

func (border *Border) Tick(event tl.Event) {
}
