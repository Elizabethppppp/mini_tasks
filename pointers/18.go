package pointers

func Reset(pp **int) {
	*pp = nil
}

func Redirect(pp **int, target *int) {
	*pp = target
}
