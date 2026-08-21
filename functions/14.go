package functions

func ForEach(s []int, f func(int) bool) {
	for _, v := range s {
		if !f(v) {
			return
		}
	}
}
