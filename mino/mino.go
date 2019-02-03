package mino

import (
	"math/rand"
	"time"
)

var (
	minos = [][][][]rune{
		{
			{
				{'.', '.', '.', '.'},
				{'.', 'G', 'G', '.'},
				{'G', 'G', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', 'G', '.', '.'},
				{'.', 'G', 'G', '.'},
				{'.', '.', 'G', '.'},
				{'.', '.', '.', '.'},
			},
		},
		{
			{
				{'.', '.', '.', '.'},
				{'.', 'R', 'R', '.'},
				{'.', '.', 'R', 'R'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', '.', '.', 'R'},
				{'.', '.', 'R', 'R'},
				{'.', '.', 'R', '.'},
				{'.', '.', '.', '.'},
			},
		},
		{
			{
				{'.', '.', '.', '.'},
				{'.', 'Y', 'Y', '.'},
				{'.', 'Y', 'Y', '.'},
				{'.', '.', '.', '.'},
			},
		},
		{
			{
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'C', 'C', 'C', 'C'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', 'C', '.', '.'},
				{'.', 'C', '.', '.'},
				{'.', 'C', '.', '.'},
				{'.', 'C', '.', '.'},
			},
		},
		{
			{
				{'.', '.', '.', '.'},
				{'.', 'M', '.', '.'},
				{'M', 'M', 'M', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'M', '.', '.', '.'},
				{'M', 'M', '.', '.'},
				{'M', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'M', 'M', 'M', '.'},
				{'.', 'M', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', '.', 'M', '.'},
				{'.', 'M', 'M', '.'},
				{'.', '.', 'M', '.'},
				{'.', '.', '.', '.'},
			},
		},
		{
			{
				{'.', '.', '.', '.'},
				{'.', 'b', '.', '.'},
				{'.', 'b', 'b', 'b'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', 'b', 'b', '.'},
				{'.', 'b', '.', '.'},
				{'.', 'b', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', 'b', 'b', 'b'},
				{'.', '.', '.', 'b'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', '.', '.', 'b'},
				{'.', '.', '.', 'b'},
				{'.', '.', 'b', 'b'},
				{'.', '.', '.', '.'},
			},
		},
		{
			{
				{'.', '.', '.', '.'},
				{'.', '.', 'm', '.'},
				{'m', 'm', 'm', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'m', '.', '.', '.'},
				{'m', '.', '.', '.'},
				{'m', 'm', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'m', 'm', 'm', '.'},
				{'m', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
			{
				{'.', 'm', 'm', '.'},
				{'.', '.', 'm', '.'},
				{'.', '.', 'm', '.'},
				{'.', '.', '.', '.'},
			},
		},
	}
)

type Block [][]rune

// Mino はテトロミノ
// 以下のようなテトロミノの例では、X=0, Y=0である。
// .............
// .AA..........
// .AA..........
// .............
// .............
type Mino struct {
	Index       int
	RotateIndex int
	// X はboardにおけるx座標
	X int
	// Y はboardにおけるy座標
	Y int
}

func NewMino() *Mino {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(minos))
	return newMino(n)
}

func newMino(n int) *Mino {
	return &Mino{Index: n}
}

func (m *Mino) RotateRight() {
	m.rotateRight(minos)
}

func (m *Mino) rotateRight(ms [][][][]rune) {
	mb := ms[m.Index]
	max := len(mb)
	ri := m.RotateIndex + 1
	if ri < max {
		m.RotateIndex = ri
		return
	}
	m.RotateIndex = 0
}

func (m *Mino) RotateLeft() {
	m.rotateLeft(minos)
}

func (m *Mino) rotateLeft(ms [][][][]rune) {
	mb := ms[m.Index]
	max := len(mb)
	ri := m.RotateIndex - 1
	if ri < 0 {
		m.RotateIndex = max - 1
		return
	}
	m.RotateIndex = ri
}

func (m *Mino) Block() Block {
	return m.block(minos)
}

func (m *Mino) block(ms [][][][]rune) Block {
	return ms[m.Index][m.RotateIndex]
}

func (m *Mino) Bottom() []int {
	return m.bottom(minos)
}

// 列単位でブロックを一番下から判定し、
// 最初に空でないブロックが見つかったインデックスを
// 配列に追加していく
func (m *Mino) bottom(ms [][][][]rune) []int {
	block := ms[m.Index][m.RotateIndex]
	rowMax := len(block)
	colMax := len(block[0])
	var bis []int
	for x := 0; x < colMax; x++ {
		var found bool
		for y := rowMax - 1; 0 <= y; y-- {
			c := block[y][x]
			if c != '.' {
				bis = append(bis, y+1)
				found = true
				break
			}
		}
		if !found {
			bis = append(bis, 0)
		}
	}
	return bis
}
