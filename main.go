package main

import (
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

func canDownMino(m *mino.Mino, b board.Board) bool {
	// boardの一番下に到達したら降下不可
	h := len(b)
	minoBottoms := m.Bottom()
	for _, btm := range minoBottoms {
		pos := m.Y + btm
		if h <= pos {
			return false
		}
	}

	// ミノの下にすでにミノが存在したら不可
	tops := b.Top()
	for x, btm := range minoBottoms {
		if btm == 0 {
			continue
		}
		top := tops[x+m.X]
		if top-1 <= btm {
			return false
		}
	}
	return true
}
