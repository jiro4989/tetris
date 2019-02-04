package util

func CopyMatrix(src [][]rune) (dst [][]rune) {
	for _, line := range src {
		tmp := make([]rune, len(line))
		copy(tmp, line)
		dst = append(dst, tmp)
	}
	return
}
