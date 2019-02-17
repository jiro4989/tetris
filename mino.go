package main

import (
	"fmt"
	"math/rand"
)

type (
	Board [][]int
	Row   []int
	Block [4][4]int
	Mino  struct {
		rotateIndex int
		minoIndex   int
		x           int
		y           int
	}
)

const (
	emptyMino = iota
	wallMino
	minoBlockWidth = 4
	boardOffset    = 3
)

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
	minos = [][]Block{
		{
			{
				{0, 0, 0, 0},
				{0, 2, 2, 0},
				{2, 2, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 2, 0, 0},
				{0, 2, 2, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 3, 3, 0},
				{0, 0, 3, 3},
				{0, 0, 0, 0},
			},
			{
				{0, 0, 0, 3},
				{0, 0, 3, 3},
				{0, 0, 3, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 4, 4, 0},
				{0, 4, 4, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{5, 5, 5, 5},
				{0, 0, 0, 0},
			},
			{
				{0, 5, 0, 0},
				{0, 5, 0, 0},
				{0, 5, 0, 0},
				{0, 5, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 6, 0, 0},
				{6, 6, 6, 0},
				{0, 0, 0, 0},
			},
			{
				{6, 0, 0, 0},
				{6, 6, 0, 0},
				{6, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{6, 6, 6, 0},
				{0, 6, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 0, 6, 0},
				{0, 6, 6, 0},
				{0, 0, 6, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 7, 0, 0},
				{0, 7, 7, 7},
				{0, 0, 0, 0},
			},
			{
				{0, 7, 7, 0},
				{0, 7, 0, 0},
				{0, 7, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 7, 7, 7},
				{0, 0, 0, 7},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 0, 0, 7},
				{0, 0, 0, 7},
				{0, 0, 7, 7},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 0, 8, 0},
				{8, 8, 8, 0},
				{0, 0, 0, 0},
			},
			{
				{8, 0, 0, 0},
				{8, 0, 0, 0},
				{8, 8, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{8, 8, 8, 0},
				{8, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 8, 8, 0},
				{0, 0, 8, 0},
				{0, 0, 8, 0},
				{0, 0, 0, 0},
			},
		},
	}
	currentBoard = CopyMatrix(initialBoard)
	displayBoard = CopyMatrix(initialBoard)
)

func newMino() Mino {
	n := rand.Intn(len(minos))
	mino := Mino{
		minoIndex: n,
		x:         4,
		y:         0,
	}
	return mino
}

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

func (b Block) isOverlap(target Block) bool {
	for y, row := range b {
		for x, cell := range row {
			targetCell := target[y][x]
			if cell != emptyMino && targetCell != emptyMino {
				return true
			}
		}
	}
	return false
}

func (m Mino) getBlock() Block {
	return minos[m.minoIndex][m.rotateIndex]
}

func (m Mino) canMoveRight(b Board) bool {
	if len(b[0]) < m.x+1+minoBlockWidth {
		return false
	}
	blk := b.fetchBlock(m.x+1, m.y)
	return !m.getBlock().isOverlap(blk)
}

func (m Mino) canMoveLeft(b Board) bool {
	if m.x-1 < 0 {
		return false
	}
	blk := b.fetchBlock(m.x-1, m.y)
	return !m.getBlock().isOverlap(blk)
}

func (m Mino) canMoveDown(b Board) bool {
	if len(b) < m.y+1+minoBlockWidth {
		return false
	}
	blk := b.fetchBlock(m.x, m.y+1)
	return !m.getBlock().isOverlap(blk)
}

func (m *Mino) moveRight() {
	m.x++
}

func (m *Mino) moveLeft() {
	m.x--
}

func (m *Mino) moveDown() {
	m.y++
}

func (m *Mino) rotateRight() {
	ri := m.rotateIndex + 1
	m.rotateIndex = ri
	if len(minos[m.minoIndex]) <= ri {
		m.rotateIndex = 0
	}
}

func (m *Mino) rotateLeft() {
	ri := m.rotateIndex - 1
	m.rotateIndex = ri
	if ri < 0 {
		m.rotateIndex = len(minos[m.minoIndex]) - 1
	}
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

func (row Row) isDeletable() bool {
	for _, c := range row {
		if c <= emptyMino {
			return false
		}
	}
	return true
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
