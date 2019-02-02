package main

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

const (
	background = `
		wwwwwwwwwwww
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wkkkkkkkkkkw
		wwwwwwwwwwww
	`
)

var (
	colorMapping = map[rune]termbox.Attribute{
		'k': termbox.ColorBlack,
		'K': termbox.ColorBlack | termbox.AttrBold,
		'r': termbox.ColorRed,
		'R': termbox.ColorRed | termbox.AttrBold,
		'g': termbox.ColorGreen,
		'G': termbox.ColorGreen | termbox.AttrBold,
		'y': termbox.ColorYellow,
		'Y': termbox.ColorYellow | termbox.AttrBold,
		'b': termbox.ColorBlue,
		'B': termbox.ColorBlue | termbox.AttrBold,
		'm': termbox.ColorMagenta,
		'M': termbox.ColorMagenta | termbox.AttrBold,
		'c': termbox.ColorCyan,
		'C': termbox.ColorCyan | termbox.AttrBold,
		'w': termbox.ColorWhite,
		'W': termbox.ColorWhite | termbox.AttrBold,
	}
)

func drawBackground() {
	const dc = termbox.ColorDefault
	lines := strings.Split(background, "\n")
	for y, line := range lines {
		line = strings.Replace(line, "\t", "", -1)
		for x, c := range line {
			color := colorMapping[c]
			termbox.SetCell(2*x-1, y, ' ', dc, color)
			termbox.SetCell(2*x, y, ' ', dc, color)
		}
	}
	termbox.Flush()
}
