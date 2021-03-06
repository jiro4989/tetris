package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchBlock(t *testing.T) {
	board := Board{
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	assert.Equal(t, Block{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}, board.fetchBlock(3, 0), "取得")
}

func TestDeleteRows(t *testing.T) {
	board := Board{
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	board.deleteRows()

	ret := Board{
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	assert.Equal(t, ret, board, "削除によってスライドする")
	board.deleteRows()
	assert.Equal(t, ret, board, "これ以上削除できない")
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

func TestSetMino(t *testing.T) {
	board := Board{
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	board.setMino(Mino{x: 3})

	assert.Equal(t, Board{
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 2, 2, 0, 1, 1, 1},
		{1, 1, 1, 2, 2, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}, board, "セットされる")
}

func TestIsDeletable(t *testing.T) {
	assert.Equal(t, true, Row{1, 1, 1, 1}.isDeletable(), "削除可")
	assert.Equal(t, false, Row{1, 0, 1, 1}.isDeletable(), "削除不可")
	assert.Equal(t, false, Row{0, 0, 0, 0}.isDeletable(), "削除不可")
}
