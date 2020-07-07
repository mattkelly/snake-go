package main

import "testing"

import tl "github.com/JoelOtter/termloop"

func TestIncreaseScore(t *testing.T) {
	scoreText = tl.NewText(0, 0, " Score: 0", tl.ColorBlack, tl.ColorBlue)
	IncreaseScore(5)
    if  scoreText.Text() != " Score: 5 " {
       t.Errorf("IncreaseScore test failed. Was expecting 5, got: %s", scoreText.Text())
    }
}