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
