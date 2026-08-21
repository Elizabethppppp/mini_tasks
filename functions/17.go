package functions

func BSearchRec(slice []int, search int) int {

	var recursive func(left, right int) int
	recursive = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := left + (right-left)/2

		if slice[mid] == search {
			return mid
		}

		if slice[mid] > search {
			return recursive(left, mid-1)
		}
		return recursive(mid+1, right)
	}
	return recursive(0, len(slice)-1)
}

func BSearchIter(slice []int, search int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == search {
			return mid
		}

		if slice[mid] > search {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
