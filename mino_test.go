package main

import (
	"fmt"
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

func TestFetchRow(t *testing.T) {
	board := Board{
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	assert.Equal(t, Row{0, 0, 0, 0}, board.fetchRow(0), "先頭行を取得")
	assert.Equal(t, Row{1, 0, 0, 0}, board.fetchRow(1), "1行目取得")
	assert.Equal(t, Row{0, 1, 0, 0}, board.fetchRow(2), "2行目取得")
	assert.Equal(t, Row{0, 0, 1, 0}, board.fetchRow(3), "3行目取得")
	assert.Equal(t, Row{0, 0, 0, 1}, board.fetchRow(4), "4行目取得")
	assert.Equal(t, Row{1, 1, 1, 1}, board.fetchRow(5), "5行目取得")
}

func TestDeleteRow(t *testing.T) {
	board := Board{
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	}

	board.deleteRow(5)

	assert.Equal(t, Board{
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1},
	}, board, "削除によってスライドする")

	for _, row := range board {
		fmt.Println(row)
	}
}
