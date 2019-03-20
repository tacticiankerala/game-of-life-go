package main

import (
	"math/rand"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	gameoflife "github.com/tacticiankerala/game-of-life/game-of-life"
	"github.com/tacticiankerala/game-of-life/gui"
)

func UniverseToScreen(universe *gameoflife.Universe, screen *gui.Screen) {
	rand.Seed(int64((universe.Width * universe.Height) + time.Now().Nanosecond()))
	for i := 0; i < universe.Height; i++ {
		padding := 0
		for j := 0; j < universe.Width; j++ {
			var ch rune
			if universeCell := universe.Cells[i][j]; universeCell == gameoflife.CellStateDead {
				ch = 128371
			} else if universeCell == gameoflife.CellStateNewCell {
				ch = 128035
			} else if universeCell == gameoflife.CellStateAlive {
				ch = 128019
			}
			screen.UpdateCellAt(i, j+padding, func(cell termbox.Cell) termbox.Cell {
				cell.Ch = ch
				return cell
			})
			padding += runewidth.RuneWidth(ch)
		}
	}
}
