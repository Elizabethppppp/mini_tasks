package functions

func SafeDiv(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func At(s []int, i int) (int, bool) {
	if i >= len(s) || i < 0 {
		return 0, false
	}
	return s[i], true
}
