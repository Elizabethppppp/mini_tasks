package slices

func Reverse(s []int, left, right int) {

	if left >= right || left < 0 || right > len(s) {
		return
	}

	for i, j := left, right-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Rotate(s []int, k int) {
	if k == len(s) {
		return
	}

	k = k % len(s)
	if k == 0 {
		return
	}

	Reverse(s, 0, k)
	Reverse(s, k, len(s))
	Reverse(s, 0, len(s))

}
