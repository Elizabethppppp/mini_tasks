package _map

func Steps(r []rune) map[[2]int]int {
	m := make(map[[2]int]int)
	x, y := 0, 0
	m[[2]int{x, y}]++

	for _, v := range r {
		switch v {
		case '←':
			x--
			m[[2]int{x, y}]++
		case '→':
			x++
			m[[2]int{x, y}]++
		case '↑':
			y++
			m[[2]int{x, y}]++
		case '↓':
			y--
			m[[2]int{x, y}]++
		}
	}

	return m
}
