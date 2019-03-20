package main

import (
	"os"
	"time"

	gameoflife "github.com/tacticiankerala/game-of-life/game-of-life"
	"github.com/tacticiankerala/game-of-life/gui"
)

func update(universe *gameoflife.Universe, screen *gui.Screen) {
	for {
		UniverseToScreen(universe, screen)
		time.Sleep(200 * time.Millisecond)
		universe.RefreshUniverse()
	}
}

func main() {
	screen := gui.InitScreen()
	universe := gameoflife.NewUniverseFromLiveCoordinates(30, 30, [][2]int{
		[2]int{0, 0},
		[2]int{0, 1},
		[2]int{1, 1},
		[2]int{1, 2},
		[2]int{2, 2},
		[2]int{2, 3},
		[2]int{3, 3},
		[2]int{3, 4},
		[2]int{3, 5},
		[2]int{3, 6},
		[2]int{3, 7},
		[2]int{3, 8},
		[2]int{3, 9},
		[2]int{4, 8},
		[2]int{4, 9},

		[2]int{27, 13},
		[2]int{27, 14},
		[2]int{27, 15},
		[2]int{27, 16},
		[2]int{27, 17},
		[2]int{27, 18},
		[2]int{27, 19},
		[2]int{27, 20},
		[2]int{27, 21},
		[2]int{27, 22},
		[2]int{28, 17},
		[2]int{28, 17},
		[2]int{28, 17},
		[2]int{28, 17},
		[2]int{28, 17},
		[2]int{28, 18},
		[2]int{28, 19},
		[2]int{28, 20},
		[2]int{28, 21},
		[2]int{28, 22},
	})
	go update(universe, screen)
	screen.DrawContinuouslyWithEventHandling()
	os.Exit(0)
}
