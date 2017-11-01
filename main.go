package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
	})

	border := NewBorder()
	snake := NewSnake()

	level.AddEntity(border)
	level.AddEntity(snake)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(15)
	game.Start()
}
