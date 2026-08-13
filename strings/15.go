package strings

func Caesar(s string, shift int) string {
	runes := []rune(s)
	shift = shift % 26
	if shift < 0 {
		shift += 26
	}

	for i, val := range runes {
		if val >= 'a' && val <= 'z' {
			base := 'a'
			runes[i] = rune((int(val)-int(base)+shift)%26 + int(base))
		}
		if val >= 'A' && val <= 'Z' {
			base := 'A'
			runes[i] = rune((int(val)-int(base)+shift)%26 + int(base))
		}
	}
	return string(runes)
}
