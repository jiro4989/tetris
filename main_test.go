package main

import (
	"testing"

	"github.com/jiro4989/tetris/board"
	"github.com/jiro4989/tetris/mino"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestCanDownMino(t *testing.T) {
	vm := mino.Mino{Index: 2}
	m := &vm
	b := board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	assert.Equal(t, true, canDownMino(m, b))

	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'A', 'A', 'A', 'A'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', 'A', '.', '.'},
		{'A', 'A', 'A', '.'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	b = board.Board{
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
	}
	assert.Equal(t, true, canDownMino(m, b))

	b = board.Board{
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
		{'A', '.', 'A', 'A'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	b = board.Board{
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
		{'A', '.', '.', 'A'},
		{'A', 'A', '.', 'A'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	vm = mino.Mino{Index: 2, X: 1, Y: 1}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	vm = mino.Mino{Index: 2, X: 1, Y: 2}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	vm = mino.Mino{Index: 2, X: -1, Y: 0}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
	}
	assert.Equal(t, true, canDownMino(m, b))

	vm = mino.Mino{}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', 'A', '.'},
	}
	assert.Equal(t, true, canDownMino(m, b))

	vm = mino.Mino{}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', 'A', '.'},
		{'.', '.', 'A', '.'},
	}
	assert.Equal(t, false, canDownMino(m, b))

	vm = mino.Mino{}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', 'A'},
		{'.', '.', '.', 'A'},
	}
	assert.Equal(t, true, canDownMino(m, b))

	vm = mino.Mino{X: 1}
	m = &vm
	b = board.Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', 'A'},
		{'.', '.', '.', 'A'},
	}
	assert.Equal(t, false, canDownMino(m, b))

}
