package main

import (
	"math/rand"
)

type (
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
)

var (
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
)

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

func newMino() Mino {
	n := rand.Intn(len(minos))
	mino := Mino{
		minoIndex: n,
		x:         4,
		y:         0,
	}
	return mino
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
