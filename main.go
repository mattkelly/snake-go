package main

import (
	"fmt"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var score = 0
var game *tl.Game
var border *Border
var scoreText *tl.Text

func IncreaseScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score))
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
	scoreText = tl.NewText(0, 0, " Score: 0", tl.ColorBlack, tl.ColorBlue)

	mainLevel.AddEntity(border)
	mainLevel.AddEntity(snake)
	mainLevel.AddEntity(food)
	mainLevel.AddEntity(scoreText)

	game.Screen().SetLevel(mainLevel)
	game.Screen().SetFps(10)
	game.Start()
}
