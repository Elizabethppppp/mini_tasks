package strings

func Sanitize(s string) string {
	result := string([]rune(s))
	return result
}
