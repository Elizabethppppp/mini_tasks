package functions

import "fmt"

func Hanoi(n int, from, to, via string) {
	if n == 0 {
		return
	}

	Hanoi(n-1, from, to, via)
	fmt.Printf("Big disk %d from %s, to %s\n", n, from, to)
	Hanoi(n-1, via, to, from)
}
