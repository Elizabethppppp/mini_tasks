package slices

func InsertAt1(s []int, i, v int) []int {
	result := make([]int, 0, len(s)+1)
	result = append(result, s[:i]...)
	result = append(result, v)
	result = append(result, s[i:]...)

	return result
}

func InsertAt2(s []int, i, v int) []int {
	if cap(s) < len(s)+1 {
		res := make([]int, len(s)+1)
		copy(res, s[:i])
		res[i] = v
		copy(res[i+1:], s[i:])
		return res
	}

	s = s[:len(s)+1]
	copy(s[i+1:], s[:len(s)+1])
	s[i] = v
	return s
}
