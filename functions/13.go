package functions

func Multiplier(k int) func(int) int {
	return func(x int) int {
		return k * x
	}
}

func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}
