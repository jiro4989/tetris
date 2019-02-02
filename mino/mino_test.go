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
	assert.Equal(t, &Mino{index: 0}, newMino(0))
	assert.Equal(t, &Mino{index: 1}, newMino(1))
}

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
	assert.Equal(t, Block{
		{'.', '.', '.', '.'},
		{'.', 'G', 'G', '.'},
		{'G', 'G', '.', '.'},
		{'.', '.', '.', '.'},
	}, m.block(ms))

	m = &Mino{index: 1, rotateIndex: 1}
	assert.Equal(t, Block{
		{'.', '.', '.', 'R'},
		{'.', '.', 'R', 'R'},
		{'.', '.', 'R', '.'},
		{'.', '.', '.', '.'},
	}, m.block(ms))
}

func TestBottom(t *testing.T) {
	m := &Mino{}
	assert.Equal(t, 3, m.bottom(ms), "正常系")

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
	assert.Equal(t, 5, m.bottom(ms2), "一番下を取得できる")

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
	assert.Equal(t, 1, m.bottom(ms2), "一番上を取得できる")

	ms2 = [][][][]rune{
		{
			{
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
			},
		},
	}
	assert.Equal(t, 0, m.bottom(ms2), "空のときは0")

}
