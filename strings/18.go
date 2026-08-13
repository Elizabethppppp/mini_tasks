package strings

func Freq(s string) map[rune]int {

	if s == "" {
		return map[rune]int{}
	}

	l := MyRuneCount(s)
	result := make(map[rune]int, l)

	for _, val := range s {
		result[val]++
	}

	return result

}
