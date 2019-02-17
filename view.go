package main

import (
	termbox "github.com/nsf/termbox-go"
)

var (
	colorMapping = map[int]termbox.Attribute{
		0: termbox.ColorWhite,
		1: termbox.ColorGreen | termbox.AttrBold,
		2: termbox.ColorRed | termbox.AttrBold,
		3: termbox.ColorCyan | termbox.AttrBold,
		4: termbox.ColorBlue | termbox.AttrBold,
		5: termbox.ColorYellow | termbox.AttrBold,
		6: termbox.ColorMagenta | termbox.AttrBold,
		7: termbox.ColorMagenta | termbox.AttrBold,
		8: termbox.ColorMagenta | termbox.AttrBold,
	}
)

func drawBackground() {
	updateDisplayBoard(currentMino)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	const dc = termbox.ColorDefault
	for y, row := range displayBoard {
		for x, cell := range row {
			color := colorMapping[cell]
			termbox.SetCell(2*x-1, y, ' ', dc, color)
			termbox.SetCell(2*x, y, ' ', dc, color)
		}
	}
	y := len(displayBoard) + 2
	const blk = termbox.ColorBlack
	const wht = termbox.ColorWhite
	setTextToDisplay("h: move left, j: move down k: move right", 0, y, blk, wht)
	setTextToDisplay("Space: move bottom", 0, y+1, blk, wht)
	setTextToDisplay("q or Esc: quit game", 0, y+2, blk, wht)
	termbox.Flush()
}

func setTextToDisplay(text string, x, y int, fg, bg termbox.Attribute) {
	for i, c := range text {
		termbox.SetCell(2*x+i, y, c, fg, bg)
		termbox.SetCell(2*x+i+1, y, ' ', fg, bg)
	}
}

func CopyMatrix(src Board) (dst Board) {
	for _, line := range src {
		tmp := make([]int, len(line))
		copy(tmp, line)
		dst = append(dst, tmp)
	}
	return
}
