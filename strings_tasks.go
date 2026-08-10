package main

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.Trim(strings.ToLower(s), " ")
	runes := []rune(s)
	flag := false

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] == runes[len(runes)-1-i] {
			flag = true
			continue
		}
	}

	return flag
}

func MyRepeat(s string, n int) string {
	if n == 0 {
		return ""
	}

	str := new(strings.Builder)

	for range n {
		str.WriteString(s)
	}

	return str.String()

}

func Capitalize(s string) string {

	r := []rune(s)

	for i, val := range r {
		if unicode.IsLetter(val) {
			r[i] = unicode.ToUpper(val)
			break
		}
	}
	return string(r)
}
