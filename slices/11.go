package slices

func Head(buf []int, n int) []int {
	b := make([]int, len(buf))
	copy(b, buf)
	return b[0:n]
}
