package _map

func GroupByLen(words []string) map[int][]string {
	if words == nil || len(words) == 0 {
		return make(map[int][]string)
	}

	m := make(map[int][]string)

	for _, word := range words {
		l := len([]rune(word))
		m[l] = append(m[l], word)
	}

	return m
}
