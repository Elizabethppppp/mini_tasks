package _map

import (
	"slices"
)

func GroupAnagrams(words []string) [][]string {
	m := make(map[string][]string)

	for _, w := range words {
		runes := []rune(w)
		slices.Sort(runes)
		key := string(runes)
		m[key] = append(m[key], w)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}
