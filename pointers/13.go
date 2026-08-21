package pointers

import "strconv"

func SetAge(a *int) {
}

func AgeString(a *int) string {
	var s string
	if a == nil {
		s="не указан"
		return s
	}

	return strconv.Itoa(*a)
}
