package functions

func Filter(s []int, keep func(int) bool) []int {
	result := make([]int, 0, len(s))
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map(s []int, f func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Reduce(s []int, init int, f func(acc, x int) int) int {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
