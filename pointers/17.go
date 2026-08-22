package pointers

type Big [4096]int

func SumVal(b Big) int {
	sum := 0
	for _, v := range b {
		sum += v
	}
	return sum
}

func SumPtr(b *Big) int {
	sum := 0
	for _, v := range *b {
		sum += v
	}
	return sum
}
