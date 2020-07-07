package main

import "testing"
//import "fmt"
//import tl "github.com/JoelOtter/termloop"

func TestRandInRange(t *testing.T) {
	test := randInRange(1, 50)
	// randInRange should never go below or above the parameters
    if  test < 1 || test > 50 {
	   t.Fail()
    }
}