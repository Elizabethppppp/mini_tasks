package main

import (
	"fmt"
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

func Invert(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func SortedKeys(m map[string]int) {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Print(k, ":", m[k], "  ")
	}
}

type Set struct {
	dataSet map[string]struct{}
}

func NewSet() *Set {
	return &Set{dataSet: make(map[string]struct{})}
}

func (s *Set) Add(key string) {
	if s.dataSet == nil {
		return
	}
	s.dataSet[key] = struct{}{}
}

func (s *Set) Del(key string) {
	delete(s.dataSet, key)
}

func (s *Set) Len() int {
	return len(s.dataSet)
}

func (s *Set) Has(key string) bool {
	_, ok := s.dataSet[key]
	return ok
}
