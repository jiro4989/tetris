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
	termbox.Flush()
}

func CopyMatrix(src Board) (dst Board) {
	for _, line := range src {
		tmp := make([]int, len(line))
		copy(tmp, line)
		dst = append(dst, tmp)
	}
	return
}
