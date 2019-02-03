package mino

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

func TestNewMino(t *testing.T) {
	assert.Equal(t, &Mino{Index: 0}, newMino(0))
	assert.Equal(t, &Mino{Index: 1}, newMino(1))
}

func TestRotateRight(t *testing.T) {
	m := &Mino{}

	m.rotateRight(ms)
	assert.Equal(t, &Mino{RotateIndex: 1}, m, "通常の右回転")

	m.rotateRight(ms)
	assert.Equal(t, &Mino{RotateIndex: 0}, m, "右回転で先頭のインデックスに戻る")

	m = &Mino{Index: 2}

	m.rotateRight(ms)
	assert.Equal(t, &Mino{Index: 2, RotateIndex: 0}, m, "回転対象がない場合のインデックスは変化しない")
}

func TestRotateLeft(t *testing.T) {
	m := &Mino{RotateIndex: 1}

	m.rotateLeft(ms)
	assert.Equal(t, &Mino{RotateIndex: 0}, m, "通常の左回転")

	m.rotateLeft(ms)
	assert.Equal(t, &Mino{RotateIndex: 1}, m, "左回転で末尾のインデックスに戻る")

	m = &Mino{Index: 2}

	m.rotateLeft(ms)
	assert.Equal(t, &Mino{Index: 2, RotateIndex: 0}, m, "回転対象がない場合のインデックスは変化しない")
}

func TestBlock(t *testing.T) {
	m := &Mino{}
	assert.Equal(t, Block{
		{'.', '.', '.', '.'},
		{'.', 'G', 'G', '.'},
		{'G', 'G', '.', '.'},
		{'.', '.', '.', '.'},
	}, m.block(ms))

	m = &Mino{Index: 1, RotateIndex: 1}
	assert.Equal(t, Block{
		{'.', '.', '.', 'R'},
		{'.', '.', 'R', 'R'},
		{'.', '.', 'R', '.'},
		{'.', '.', '.', '.'},
	}, m.block(ms))
}

func TestBottom(t *testing.T) {
	m := &Mino{}
	assert.Equal(t, []int{3, 3, 2, 0}, m.bottom(ms), "正常系")

	ms2 := [][][][]rune{
		{
			{
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'G', 'G', 'G', 'G'},
			},
		},
	}
	assert.Equal(t, []int{5, 5, 5, 5}, m.bottom(ms2), "一番下を取得できる")

	ms2 = [][][][]rune{
		{
			{
				{'G', 'G', 'G', 'G'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
		},
	}
	assert.Equal(t, []int{1, 1, 1, 1}, m.bottom(ms2), "一番上を取得できる")

	ms2 = [][][][]rune{
		{
			{
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
		},
	}
	assert.Equal(t, []int{0, 0, 0, 0}, m.bottom(ms2), "空のときは0")

}
