package pointers

func NewCounter() *int {
	n := 0
	return &n
}

func local() int { n := 0; return n }
