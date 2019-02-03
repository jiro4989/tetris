package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestCanDownMino(t *testing.T) {
	// m := mino.NewMino()
	// b := board.Board{
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// }
	// assert.Equal(t, true, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// }
	// assert.Equal(t, false, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// }
	// assert.Equal(t, false, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// 	{'A', 'A', 'A', 'A'},
	// }
	// assert.Equal(t, false, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'.', '.', '.', '.'},
	// 	{'.', '.', '.', '.'},
	// 	{'.', 'A', '.', '.'},
	// 	{'A', 'A', 'A', '.'},
	// }
	// assert.Equal(t, false, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// }
	// assert.Equal(t, true, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', 'A', 'A'},
	// }
	// assert.Equal(t, false, canDownMino(m, b))
	//
	// b = board.Board{
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// 	{'A', '.', '.', 'A'},
	// 	{'A', 'A', '.', 'A'},
	// }
	// assert.Equal(t, false, canDownMino(m, b))

}
