package _map

func MapsEqual(a, b map[string]int) bool {

	if len(a) != len(b) {
		return false
	}

	if (a == nil) && (b == nil) {
		return true
	}


	for k, v := range b {
		if a[k] != v {
			return false
		}
	}

	return true
}
