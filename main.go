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
var isFullscreen *bool
var isRealFood *bool

type endgameScreen struct {
	*tl.BaseLevel
}

// Handle events on the endLevel. Enter resets.
func (eg *endgameScreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		if event.Key == tl.KeyEnter {
			score = 0
			game.Screen().SetLevel(newMainLevel(isFullscreen))
		}
	}
}

// IncreaseScore increases the score by the given amount and updates the
// score text.
func IncreaseScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score, " "))
}

// EndGame should be called when the game ends due to e.g. dying.
func EndGame() {
	endLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorRed,
	})
	el := new(endgameScreen)
	el.BaseLevel = endLevel
	var PromptQuestion, PromptText *tl.Text
	PromptQuestion = tl.NewText(34, 17, " Play Again? ", tl.ColorBlue, tl.ColorWhite)
	PromptText = tl.NewText(34, 18, " Press Enter ", tl.ColorBlue, tl.ColorWhite)
	scoreText.SetPosition(35, 14)
	scoreText.SetColor(tl.ColorBlue, tl.ColorWhite)
	el.AddEntity(scoreText)
	el.AddEntity(PromptQuestion)
	el.AddEntity(PromptText)

	game.Screen().SetLevel(el)
}

func newMainLevel(isFullscreen *bool) tl.Level{

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
	return mainLevel
}

func main() {
	isFullscreen = flag.Bool("fullscreen", false, "Play fullscreen!")
	isRealFood = flag.Bool("realfood", false, "Realistic food options")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()

	mainLevel := newMainLevel(isFullscreen)

	game.Screen().SetLevel(mainLevel)
	game.Screen().SetFps(10)
	game.Start()
}
