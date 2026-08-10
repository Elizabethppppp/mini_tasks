package main

import "fmt"

func DivMod(a, b int) (int, int) {
	return a / b, a % b
}

func MinMax(s []int) (int, int) {
	current := s[0]
	for i := range s {
		if s[i] < current {
			current = s[i]
		}
	}
	currentMax := s[0]
	for i := range s {
		if s[i] > currentMax {
			currentMax = s[i]
		}
	}
	return current, current

}

func SumFunction(nums ...int) int {
	fmt.Printf("%T ", nums)

	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

type Op func(int, int) int

func AddType(a, b int) int {
	return a + b
}

func MultipleType(a, b int) int {
	return a * b
}

func SubstractType(a, b int) int {
	return a - b

}

func ModType(a, b int) int {
	return a % b
}

func Fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Fact(n-1)
}

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
