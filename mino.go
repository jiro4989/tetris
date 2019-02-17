package main

type (
	Board     [][]int8
	MinoBoard struct {
		board  Board
		offset uint8
	}
	Row   []int8
	Block [4][4]int8
	Mino  struct {
		rotateIndex int8
		minoIndex   int8
		x           int
		y           int
	}
)

const (
	minoBlockWidth = 4
	emptyMino      = iota
	wallMino
)

var (
	initialBoard = [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	minos = [][][][]int{
		{
			{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{1, 1, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 1, 1},
				{0, 0, 0, 0},
			},
			{
				{0, 0, 0, 1},
				{0, 0, 1, 1},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 1, 1, 1},
				{0, 0, 0, 0},
			},
			{
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 0, 0, 0},
			},
			{
				{1, 0, 0, 0},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 1, 1},
				{0, 0, 0, 0},
			},
			{
				{0, 1, 1, 0},
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 1, 1, 1},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 0, 0, 1},
				{0, 0, 0, 1},
				{0, 0, 1, 1},
				{0, 0, 0, 0},
			},
		},
		{
			{
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{1, 1, 1, 0},
				{0, 0, 0, 0},
			},
			{
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{1, 1, 1, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			{
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
		},
	}
	currentBoard = initialBoard
	displayBoard = initialBoard
)

func (b Board) fetchBlock(x, y int) (ret Block) {
	var i int
	for y2 := y; y2 < y+minoBlockWidth; y2++ {
		var j int
		for x2 := y; x2 < y+minoBlockWidth; x2++ {
			cell := b[y2][x2]
			ret[i][j] = cell
			j++
		}
		i++
	}
	return
}

func (b Block) isOverlap(target Block) bool {
	for y, row := range b {
		for x, cell := range row {
			targetCell := b[y][x]
			overlap := cell + targetCell
			if 1 < overlap {
				return true
			}
		}
	}
	return false
}

func (m Mino) getBlock() Block {
	return minos[m.minoIndex][m.rotateIndex]
}

func (m Mino) canMoveRight(b Board) bool {
	if len(b[0]) < m.x+1+minoBlockWidth {
		return false
	}
	blk := b.fetchBlock(m.x+1, ym.y)
	return !m.getBlock().isOverlap(blk)
}

func (m Mino) canMoveLeft(b Board) bool {
	if m.x-1 < 0 {
		return false
	}
	blk := b.fetchBlock(m.x-1, ym.y)
	return !m.getBlock().isOverlap(blk)
}

func (m Mino) canMoveDown(b Board) bool {
	if len(b) < m.y+1+minoBlockWidth {
		return false
	}
	blk := b.fetchBlock(m.x, m.y+1)
	return !m.getBlock().isOverlap(blk)
}

func (m Mino) moveRight() {
	m.x++
}

func (m Mino) moveLeft() {
	m.x--
}

func (m Mino) moveDown() {
	m.y++
}

func (m Mino) rotateRight() {
	ri := m.rotateIndex + 1
	m.rotateIndex = ri
	if len(minos[m.minoIndex]) <= ri {
		m.rotateIndex = 0
	}
}

func (m Mino) rotateLeft() {
	ri := m.rotateIndex - 1
	m.rotateIndex = ri
	if ri < 0 {
		m.rotateIndex = len(minos[m.minoIndex]) - 1
	}
}

func (row Row) isDeletable() bool {
	for _, c := range row {
		if c <= 0 {
			return false
		}
	}
	return true
}

func (mb MinoBoard) fetchRow(n int) Row {
	row := mb.board[n]
	return row[mb.offset : len(row)-1-mb.offset]
}

func (b Board) setMino(m Mino) {
	blk := m.getBlock()
	for y, row := range blk {
		for x, cell := range row {
			if cell != emptyMino {
				b[y+m.y][x+m.x] = cell
			}
		}
	}
}

func updateDisplayBoard(m Mino) {
	displayBoard = currentBoard
	displayBoard.setMino(m)
}

func updateCurrentBoard(m Mino) {
	currentBoard.setMino(m)
	displayBoard = currentBoard
}
