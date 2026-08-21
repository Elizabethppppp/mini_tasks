package functions

import "slices"

func MySortFunc(str []string) {

	slices.SortFunc(str, func(a, b string) int {
		la, lb := len([]rune(a)), len([]rune(b))

		if la != lb {
			return la - lb
		}

		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
}
