package strings

func MyRuneCount(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i]&0xC0 != 0x80 {
			count++
		}
	}
	return count
}
