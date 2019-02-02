package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ms = [][][][]rune{
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
	}
)

func TestRotateRight(t *testing.T) {
	m := &Mino{}

	m.rotateRight(ms)
	assert.Equal(t, &Mino{rotateIndex: 1}, m, "通常の右回転")

	m.rotateRight(ms)
	assert.Equal(t, &Mino{rotateIndex: 0}, m, "右回転で先頭のインデックスに戻る")

	m = &Mino{index: 2}

	m.rotateRight(ms)
	assert.Equal(t, &Mino{index: 2, rotateIndex: 0}, m, "回転対象がない場合のインデックスは変化しない")
}

func TestRotateLeft(t *testing.T) {
	m := &Mino{rotateIndex: 1}

	m.rotateLeft(ms)
	assert.Equal(t, &Mino{rotateIndex: 0}, m, "通常の左回転")

	m.rotateLeft(ms)
	assert.Equal(t, &Mino{rotateIndex: 1}, m, "左回転で末尾のインデックスに戻る")

	m = &Mino{index: 2}

	m.rotateLeft(ms)
	assert.Equal(t, &Mino{index: 2, rotateIndex: 0}, m, "回転対象がない場合のインデックスは変化しない")
}
func TestBlock(t *testing.T) {
	m := &Mino{}
	assert.Equal(t, [][]rune{
		{'.', '.', '.', '.'},
		{'.', 'G', 'G', '.'},
		{'G', 'G', '.', '.'},
		{'.', '.', '.', '.'},
	}, m.block(ms))

	m = &Mino{index: 1, rotateIndex: 1}
	assert.Equal(t, [][]rune{
		{'.', '.', '.', 'R'},
		{'.', '.', 'R', 'R'},
		{'.', '.', 'R', '.'},
		{'.', '.', '.', '.'},
	}, m.block(ms))

}

func TestNewMino(t *testing.T) {
	assert.Equal(t, &Mino{index: 0}, newMino(0))
	assert.Equal(t, &Mino{index: 1}, newMino(1))
}
