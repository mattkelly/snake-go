package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
	})

	// TODO dynamically size based on screen width
	border := NewBorder(50, 20)
	snake := NewSnake()

	level.AddEntity(border)
	level.AddEntity(snake)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(10)
	game.Start()
}
