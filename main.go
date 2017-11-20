package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
	"github.com/nsf/termbox-go"
)

var score = 0
var game *tl.Game
var border *Border
var scoreText *tl.Text

// IncreaseScore increases the score by the given amount and updates the
// score text.
func IncreaseScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score))
}

// EndGame should be called when the game ends due to e.g. dying.
func EndGame() {
	endLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorRed,
	})

	game.Screen().SetLevel(endLevel)
}

func main() {
	isFullscreen := flag.Bool("fullscreen", false, "Play fullscreen!")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()

	mainLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	width, height := 80, 30
	if *isFullscreen {
		// Must initialize Termbox before getting the terminal size
		termbox.Init()
		width, height = termbox.Size()
	}
	border = NewBorder(width, height)

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
