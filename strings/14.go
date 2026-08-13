package strings

import (
	"strings"
	"unicode"
)

func MyFields(s string) []string {
	if s == "" || s == " " {
		return []string{}
	}

	var results []string
	var word strings.Builder

	for _, val := range s {
		if unicode.IsSpace(val) {
			if word.Len() > 0 {
				results = append(results, word.String())
				word.Reset()
			}
		} else {
			word.WriteRune(val)
		}
	}
	if word.Len() > 0 {
		results = append(results, word.String())
	}
	return results
}
