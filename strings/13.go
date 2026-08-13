package strings

func MyCount(s, sub string) int {

	var count int

	if s == "" || sub == "" {
		return 0
	}

	for i := 0; i < len(s)-len(sub)+1; {
		if s[i:i+len(sub)] == sub {
			count++
			i += len(sub)
		} else {
			i++
		}
	}

	return count
}
