package slices

func FilterEven(s []int) []int {
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			s[index] = s[i]
			index++
		}
	}
	return s[:index]
}
