package main

func SafeDeref(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func IncVal(x int) {
	x += 1
}
func IncPtr(p *int) {
	*p += 1
}
