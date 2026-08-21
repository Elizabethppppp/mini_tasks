package pointers

import (
	"strconv"
	"strings"
)

func ParsePoint(s string) (x, y int, ok bool) {
	str := strings.Split(s, ";")
	if len(str) != 2 {
		return 0, 0, false
	}
	x, err1 := strconv.Atoi(str[0])
	if err1 != nil {
		return 0, 0, false
	}
	y, err2 := strconv.Atoi(str[1])
	if err2 != nil {
		return 0, 0, false
	}

	return x, y, true

}

func ParsePointInto(s string, x, y *int) bool {
	if x == nil || y == nil {
		return false
	}

	str := strings.Split(s, ";")
	if len(str) != 2 {
		return false
	}
	xVal, err1 := strconv.Atoi(str[0])
	if err1 != nil {
		return false
	}
	yVal, err2 := strconv.Atoi(str[1])
	if err2 != nil {
		return false
	}

	*x=xVal
	*y=yVal
	return true
}
