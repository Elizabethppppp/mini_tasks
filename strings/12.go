package strings

func ReverseRunes(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}
	return string(r)
}
