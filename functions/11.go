package functions

func Fib() func(int) int {
	cashes := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if v, ok := cashes[n]; ok {
			return v
		}
		result := fib(n-1) + fib(n-2)
		cashes[n] = result
		return result
	}
	return fib
}
