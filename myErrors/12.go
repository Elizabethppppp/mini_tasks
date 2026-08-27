package myErrors

func check(n int) error {
	if n < 0 {
		return &OutOfRange{n, 0, 150}
	}
	return nil
}

func validate(n int) error {
	err := check(n)
	if err != nil {
		return err
	}
	return nil
}
