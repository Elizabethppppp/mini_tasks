package slices

import "slices"

func Copy1(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}

func Copy2(in []int) []int {
	clone := append([]int(nil), in...)
	return clone
}

func Copy3(in []int) []int {
	clone := slices.Clone(in)
	return clone

}
