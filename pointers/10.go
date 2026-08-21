package pointers

import "fmt"

func ModInt(x int) {

	x = 100
}

func ModSlice(s []int) {

	s[0] = 100
	fmt.Println(s)
	s = append(s, 1)
}

func ModMap(m map[string]int) {
	m["new"] = 1
}

func ModArr(a [3]int) {
	a[0] = 100
}
