package _map

func MergeSum(a, b map[string]int) map[string]int {
	switch {
	case a == nil:
		return b
	case b == nil:
		return a
	case a == nil && b == nil:
		return nil
	case len(a) == 0:
		return b
	case len(b) == 0:
		return a
	}

	m := make(map[string]int, len(a)+len(b))
	for k, v := range a {
		m[k] = v
	}
	for k, v := range b {
		m[k] += v
	}
	return m
}
