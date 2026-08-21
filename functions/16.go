package functions

type Pred func(int) bool

func And(a, b Pred) Pred {
	return func(i int) bool {
		return a(i) && b(i)
	}
}

func Or(a, b Pred) Pred {
	return func(i int) bool {
		return a(i) || b(i)
	}
}

func Not(a Pred) Pred {
	return func(i int) bool {
		return !a(i)
	}
}
