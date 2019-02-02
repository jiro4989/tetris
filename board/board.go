package board

type Board [][]rune

func (b Board) Top() int {
	for i, line := range b {
		for _, c := range line {
			if c != '.' {
				return i + 1
			}
		}
	}
	return 0
}
