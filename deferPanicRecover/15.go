package deferPanicRecover

import "strconv"

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("invalid integer literal: " + s)
	}
	return n
}
