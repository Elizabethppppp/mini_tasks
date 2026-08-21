package pointers

func GetOr(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}
