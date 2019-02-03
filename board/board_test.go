package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMino(t *testing.T) {
	assert.Equal(t, 1, Board{{'A'}, {'.'}, {'.'}, {'.'}}.Top())
	assert.Equal(t, 3, Board{{'.'}, {'.'}, {'A'}, {'.'}}.Top())
	assert.Equal(t, 4, Board{{'.'}, {'.'}, {'.'}, {'A'}}.Top())
	assert.Equal(t, 5, Board{{'.'}, {'.'}, {'.'}, {'.'}}.Top())
}
