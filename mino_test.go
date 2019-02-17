package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOverlap(t *testing.T) {
	b1 := Block{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}

	assert.Equal(t, false, b1.isOverlap(Block{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}), "空ブロックとの判定ではfalse")

	assert.Equal(t, false, b1.isOverlap(Block{
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 0, 0, 1},
		{1, 1, 1, 1},
	}), "ぎりぎり重ならない状態ならfalse")

	assert.Equal(t, true, b1.isOverlap(Block{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}), "一箇所だけ重なるときはtrue")

	assert.Equal(t, true, b1.isOverlap(Block{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 0},
	}), "一箇所だけ重なるときはtrue")

	assert.Equal(t, true, b1.isOverlap(Block{
		{0, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}), "全部重なるときはtrue")

	assert.Equal(t, true, b1.isOverlap(Block{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}), "全部ブロックが埋まってるときはtrue")

}

func TestGetBlock(t *testing.T) {
	mino := Mino{minoIndex: 0}
	assert.Equal(t, Block{
		{0, 0, 0, 0},
		{0, 2, 2, 0},
		{2, 2, 0, 0},
		{0, 0, 0, 0},
	}, mino.getBlock())
}

var board = Board{
	{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func TestCanMoveRight(t *testing.T) {
	mino := Mino{minoIndex: 0, x: 3, y: 0}
	assert.Equal(t, true, mino.canMoveRight(board))
	mino.x++
	assert.Equal(t, false, mino.canMoveRight(board))
	mino.x = 111
	assert.Equal(t, false, mino.canMoveRight(board))
}

func TestCanMoveLeft(t *testing.T) {
	mino := Mino{minoIndex: 0, x: 3, y: 0}
	assert.Equal(t, false, mino.canMoveLeft(board))
	mino.x++
	assert.Equal(t, true, mino.canMoveLeft(board))
	mino.x = -1
	assert.Equal(t, false, mino.canMoveLeft(board))
}

func TestCanMoveDown(t *testing.T) {
	mino := Mino{minoIndex: 0, x: 3, y: 0}
	assert.Equal(t, true, mino.canMoveDown(board))
	mino.y = 1
	assert.Equal(t, false, mino.canMoveDown(board))
	mino.y = 3
	assert.Equal(t, false, mino.canMoveDown(board))
	mino.y = 111
	assert.Equal(t, false, mino.canMoveDown(board))
}

func TestMove(t *testing.T) {
	mino := Mino{x: 0, y: 0}

	mino.moveRight()
	assert.Equal(t, 1, mino.x)

	mino.moveLeft()
	assert.Equal(t, 0, mino.x)

	mino.moveDown()
	assert.Equal(t, 1, mino.y)
}

func TestRotateRight(t *testing.T) {
	mino := Mino{}

	mino.rotateRight()
	assert.Equal(t, 1, mino.rotateIndex)

	mino.rotateRight()
	assert.Equal(t, 0, mino.rotateIndex)
}

func TestRotateLeft(t *testing.T) {
	mino := Mino{}

	mino.rotateLeft()
	assert.Equal(t, 1, mino.rotateIndex)

	mino.rotateLeft()
	assert.Equal(t, 0, mino.rotateIndex)
}
