package gui

import (
	"sync"
	"time"

	"github.com/nsf/termbox-go"
)

type ControlFlag struct {
	resetScreen bool
	closeScreen bool
}

type Screen struct {
	width, height int
	cells         []termbox.Cell
	mutex         sync.Mutex
	controlFlag   ControlFlag
}

func InitScreen() *Screen {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	width, height := termbox.Size()
	cells := make([]termbox.Cell, width*height)

	return &Screen{width: width, height: height, cells: cells}
}

func (screen *Screen) DrawContinuouslyWithEventHandling() {
	go screen.startEventListener()
	for {
		if screen.getResetScreenFlag() {
			screen.resetScreen()
			screen.updateResetScreenFlag(false)
		}
		if screen.getCloseScreenFlag() {
			closeScreen()
			screen.updateCloseScreenFlag(false)
			return
		}
		screen.drawOnce()
		time.Sleep(16 * time.Millisecond)
	}
}

func (screen *Screen) UpdateCellAt(x, y int, updateFunc func(termbox.Cell) termbox.Cell) {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()

	i := (x * screen.width) + y
	cell := screen.cells[i]
	updatedCell := updateFunc(cell)
	screen.cells[i].Ch = updatedCell.Ch
	screen.cells[i].Fg = updatedCell.Fg
	screen.cells[i].Bg = updatedCell.Bg
}

func closeScreen() {
	termbox.Close()
}

func (screen *Screen) resetScreen() {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()

	width, height := termbox.Size()
	cells := make([]termbox.Cell, width*height)
	screen.width = width
	screen.height = height
	screen.cells = cells
}

func (screen *Screen) getResetScreenFlag() bool {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()
	return screen.controlFlag.resetScreen
}

func (screen *Screen) updateResetScreenFlag(flag bool) {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()
	screen.controlFlag.resetScreen = flag
}

func (screen *Screen) getCloseScreenFlag() bool {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()
	return screen.controlFlag.closeScreen
}

func (screen *Screen) updateCloseScreenFlag(flag bool) {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()
	screen.controlFlag.closeScreen = flag
}

func (screen *Screen) drawOnce() {
	screen.mutex.Lock()
	defer screen.mutex.Unlock()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	copy(termbox.CellBuffer(), screen.cells)
	termbox.Flush()
}

func (screen *Screen) startEventListener() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				screen.updateCloseScreenFlag(true)
				return
			}
		case termbox.EventResize:
			screen.updateResetScreenFlag(true)
		}
	}
}
