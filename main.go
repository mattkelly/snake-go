package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var score = 0
var level *tl.BaseLevel

func main() {
	rand.Seed(time.Now().UnixNano())
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	border := NewBorder()
	snake := NewSnake()
	food := NewFood()

	level.AddEntity(border)
	level.AddEntity(snake)
	level.AddEntity(food)

	game.Screen().SetLevel(level)
	game.Screen().SetFps(10)
	game.Start()
}
