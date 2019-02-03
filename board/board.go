package board

type Board [][]rune

// 列単位でブロックを一番上から判定し、
// 最初に空でないブロックが見つかったインデックスを
// 配列に追加していく
func (b Board) Top() []int {
	rowMax := len(b)
	colMax := len(b[0])
	var bis []int
	for x := 0; x < colMax; x++ {
		var found bool
		for y := 0; y < rowMax; y++ {
			c := b[y][x]
			if c != '.' {
				bis = append(bis, y+1)
				found = true
				break
			}
		}
		if !found {
			bis = append(bis, rowMax+1)
		}
	}
	return bis
}
