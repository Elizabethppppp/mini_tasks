package _map

import "strings"

func Index(text string) map[string][]int {

	str := strings.Fields(text)
	m := make(map[string][]int, len(str))

	for i, word := range str {
		m[word] = append(m[word], i)
	}

	return m
}
