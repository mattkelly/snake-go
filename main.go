package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var score = 0
var game *tl.Game
var border *Border

func IncreaseScore(amount int) {
	score += amount
}

func EndGame() {
	endLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorRed,
	})

	game.Screen().SetLevel(endLevel)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()

	mainLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	border = NewBorder(80, 30)

	snake := NewSnake()
	food := NewFood()

	mainLevel.AddEntity(border)
	mainLevel.AddEntity(snake)
	mainLevel.AddEntity(food)

	game.Screen().SetLevel(mainLevel)
	game.Screen().SetFps(10)
	game.Start()
}
