package main

import "fmt"

func Sum(s []int) int {

	sum := 0

	for i := 0; i < len(s); i++ {
		sum += s[i]

	}
	return sum
}

func Min(a []int) int {
	min := a[0]

	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}

	return min
}

func Max(a []int) int {

	max := a[0]

	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func Multiplication1(k, n int) []int {
	m := make([]int, n)

	for i := range n {
		m[i] = k * (i + 1)
	}
	return m
}

func Multiplication2(k, n int) []int {
	m := make([]int, 0, n)

	for i := range n {
		m = append(m, k*(i+1))
	}
	return m
}

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Append(s []int) []int {
	for i := 0; i < 5; i++ {
		s = append(s, i+1)
		fmt.Println(len(s), cap(s))
	}
	return s
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
