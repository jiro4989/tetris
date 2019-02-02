package main

import (
	"fmt"

	"github.com/jiro4989/tetris/board"
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

func canDownMino(m mino.Mino, b board.Board) bool {
	// boardの一番下に到達したら降下不可
	w, h := len(b[0]), len(b)
	fmt.Println(w)
	btm := m.Bottom()
	if h <= m.Y+btm {
		return false
	}
	// ミノの下にすでにミノが存在したら不可
	return true
}
