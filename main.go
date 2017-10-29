package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorRed,
		Ch: '.',
	})

	snake := NewSnake()

	level.AddEntity(snake)

	game.Screen().SetLevel(level)
	game.Start()
}
