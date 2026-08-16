package _map

import (
	"pasha_tasks/strings"
	"sort"
)

func CountWords(s string) map[string]int {
	if s == "" {
		return map[string]int{}
	}
	m := make(map[string]int, len(s))
	words := strings.MyFields(s)

	for _, word := range words {
		m[word]++
	}
	return m
}

func TopFreq(m map[string]int) []string {

	words := make([]string, 0, len(m))
	for word := range m {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if m[words[i]] != m[words[j]] {
			return m[words[i]] > m[words[j]]
		}
		return words[i] < words[j]
	})

	if 5 > len(words) {
		return words
	}

	return words[:5]
}
