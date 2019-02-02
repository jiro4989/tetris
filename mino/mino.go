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

type Mino struct {
	index       int
	rotateIndex int
	X           int
	Y           int
}

func NewMino() *Mino {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(minos))
	return newMino(n)
}

func newMino(n int) *Mino {
	return &Mino{index: n}
}

func (m *Mino) RotateRight() {
	m.rotateRight(minos)
}

func (m *Mino) rotateRight(ms [][][][]rune) {
	mb := ms[m.index]
	max := len(mb)
	ri := m.rotateIndex + 1
	if ri < max {
		m.rotateIndex = ri
		return
	}
	m.rotateIndex = 0
}

func (m *Mino) RotateLeft() {
	m.rotateLeft(minos)
}

func (m *Mino) rotateLeft(ms [][][][]rune) {
	mb := ms[m.index]
	max := len(mb)
	ri := m.rotateIndex - 1
	if ri < 0 {
		m.rotateIndex = max - 1
		return
	}
	m.rotateIndex = ri
}

func (m *Mino) Block() Block {
	return m.block(minos)
}

func (m *Mino) block(ms [][][][]rune) Block {
	return ms[m.index][m.rotateIndex]
}

func (m *Mino) Bottom() int {
	return m.bottom(minos)
}

func (m *Mino) bottom(ms [][][][]rune) int {
	block := ms[m.index][m.rotateIndex]
	for i := len(block) - 1; 0 <= i; i-- {
		line := block[i]
		for _, c := range line {
			if c != '.' {
				return i + 1
			}
		}
	}
	return 0
}
