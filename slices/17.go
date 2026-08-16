package slices

import "fmt"

func Elements() {
	var s []int
	var prev int

	for i := 0; i < 2000; i++ {
		s = append(s, i)
		current := cap(s)
		if current != prev {
			fmt.Printf("i: %d, cap: %d\n", i, current)
			prev = current
		}
	}
}

func ElementsMake() {
	s := make([]int, 0, 2000)
	prev := cap(s)
	for i := 0; i < 2000; i++ {
		s = append(s, i)
		current := cap(s)
		if current != prev {
			fmt.Printf("i: %d, cap: %d\n", i, current)
			prev = current
		}
	}
}
