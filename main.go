package main

import (
	"fmt"
	"time"

	"github.com/jiro4989/tetris/board"
	"github.com/jiro4989/tetris/mino"
)

func main() {
	// // termboxの初期化
	// if err := termbox.Init(); err != nil {
	// 	panic(err)
	// }
	// defer termbox.Close()

	// termbox.SetInputMode(termbox.InputEsc)
	// termbox.Flush()

	// drawBackground()

	bufBoard := board.Board{
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.'},
	}
	currentBoard := copyMatrix(bufBoard)

	m := mino.NewMino()
	for {
		bufBoard = copyMatrix(currentBoard)
		blk := m.Block()
		for y, line := range blk {
			for x, c := range line {
				if c != '.' {
					bufBoard[y+m.Y][x] = c
				}
			}
		}
		if canDownMino(m, currentBoard) {
			fmt.Println("can")
			m.Y++
		} else {
			fmt.Println("not")
			currentBoard = bufBoard
		}
		for _, line := range bufBoard {
			fmt.Println(string(line))
		}
		time.Sleep(1 * time.Second)
	}

	waitKeyInput()
}

func copyMatrix(src [][]rune) (dst [][]rune) {
	for _, line := range src {
		tmp := make([]rune, len(line))
		copy(tmp, line)
		dst = append(dst, tmp)
	}
	return
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
