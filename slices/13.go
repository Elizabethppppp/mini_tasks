package slices

func DeleteAt1(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func DeleteAt2(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
