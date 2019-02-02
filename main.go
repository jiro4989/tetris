package main

import (
	"github.com/jiro4989/tetris/mino"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	// termboxの初期化
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	drawBackground()

	waitKeyInput()
}

var board = [][]rune{
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.'},
}

type Board [][]rune

func canDownMino(m mino.Mino, b Board) bool {
	// boardの一番下に到達したら降下不可
	// ミノの下にすでにミノが存在したら不可
	return true
}
