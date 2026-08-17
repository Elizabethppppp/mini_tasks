package _map

func ContainsSlice(s []int, x int) bool {
	for _, n := range s {
		if n == x {
			return true
		}
	}
	return false
}

func ContainsMap(m map[int]struct{}, x int) bool {
	_, ok := m[x]
	return ok
}
