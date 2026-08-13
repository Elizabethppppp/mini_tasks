package strings

import "unicode/utf8"

func TruncateBytes(s string, max int) string {

	if max < 3 {
		return ""
	}

	if max == 3 {
		return "..."
	}

	if len(s) <= max {
		return s
	}

	bytes := max - 3

	end := bytes
	if end > 0 && !utf8.RuneStart(s[end]) {
		end--
	}

	return string(s[:end]) + "..."

}
