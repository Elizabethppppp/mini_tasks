package slices

func InsertAt1(s []int, i, v int) []int {
	return append(append(s[:i], v), s[i:]...)
}
