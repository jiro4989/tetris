package main

import "fmt"

type (
	Board [][]int
	Row   []int
)

const boardOffset = 3

var (
	initialBoard Board = Board{
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	currentBoard = CopyMatrix(initialBoard)
	displayBoard = CopyMatrix(initialBoard)
)

func (b Board) fetchBlock(x, y int) (ret Block) {
	var i int
	for y2 := y; y2 < y+minoBlockWidth; y2++ {
		var j int
		for x2 := x; x2 < x+minoBlockWidth; x2++ {
			cell := b[y2][x2]
			ret[i][j] = cell
			j++
		}
		i++
	}
	return
}

func (b *Board) deleteRows() {
	for i := 0; i < len(*b)-boardOffset; i++ {
		row := b.fetchRow(i)
		if row.isDeletable() {
			b.deleteRow(i)
		}
	}
}

// deleteRow はレコードを削除して上のレコードをスライドしてくる
func (b *Board) deleteRow(n int) {
	for i := n; 0 < i; i-- {
		copy((*b)[i], (*b)[i-1])
	}
	copy((*b)[0], initialBoard[0])
}

func (b Board) fetchRow(n int) Row {
	row := b[n]
	return row[boardOffset : len(row)-boardOffset]
}

func (b *Board) setMino(m Mino) {
	blk := m.getBlock()
	for y, row := range blk {
		for x, cell := range row {
			if cell != emptyMino {
				(*b)[y+m.y][x+m.x] = cell
			}
		}
	}
}

func updateDisplayBoard(m Mino) {
	displayBoard = CopyMatrix(currentBoard)
	displayBoard.setMino(m)
}

func updateCurrentBoard(m Mino) {
	currentBoard.setMino(m)
	displayBoard = CopyMatrix(currentBoard)
}

// show はデバッグ用の関数なのでメインロジックにはあんまり関係ない
// 機能としてはボードをボードっぽく表示してくれてハッピーになれる感じのアレ
func (b Board) show() {
	for _, row := range b {
		var line string
		for _, c := range row {
			line += fmt.Sprintf("%d", c)
		}
		fmt.Println(line)
	}
}

func (row Row) isDeletable() bool {
	for _, c := range row {
		if c <= emptyMino {
			return false
		}
	}
	return true
}

func CopyMatrix(src Board) (dst Board) {
	for _, line := range src {
		tmp := make([]int, len(line))
		copy(tmp, line)
		dst = append(dst, tmp)
	}
	return
}
