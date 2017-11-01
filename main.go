package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	border := NewBorder()
	snake := NewSnake()
	food := NewFood(Coord{20, 20})

	level.AddEntity(border)
	level.AddEntity(snake)
	level.AddEntity(food)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(15)
	game.Start()
}
