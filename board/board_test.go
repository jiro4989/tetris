package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMino(t *testing.T) {
	assert.Equal(t, []int{4, 3, 2, 1}, Board{
		{'.', '.', '.', 'A'},
		{'.', '.', 'A', 'A'},
		{'.', 'A', 'A', '.'},
		{'A', 'A', 'A', '.'}}.Top())

	assert.Equal(t, []int{1, 3, 2, 5}, Board{
		{'A', '.', '.', '.'},
		{'.', '.', 'A', '.'},
		{'.', 'A', 'A', '.'},
		{'A', 'A', 'A', '.'}}.Top())

	assert.Equal(t, []int{5, 5, 5, 5}, Board{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'}}.Top())

}
