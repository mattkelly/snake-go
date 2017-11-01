package main

import tl "github.com/JoelOtter/termloop"

type Food struct {
	*tl.Entity
	coord Coord
}

func NewFood(coord Coord) *Food {
	f := new(Food)
	f.Entity = tl.NewEntity(5, 5, 1, 1)
	f.coord = coord
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
