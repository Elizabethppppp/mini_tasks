package myErrors

import "fmt"

func Retry(attempts int, f func() error) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		err := f()
		if err == nil {
			return nil
		}
		lastErr = err
	}

	return fmt.Errorf("после %d попыток: %w", attempts, lastErr)
}
