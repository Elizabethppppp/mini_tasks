package myErrors

import (
	"errors"
	"fmt"
)

var ErrTemporary = errors.New("temporary error")

func IsTemporary(err error) bool {
	return errors.Is(err, ErrTemporary)
}

func NewRetry(attempts int, f func() error) error {
	var lastErr error
	for i := 0; i < attempts; i++ {
		err := f()
		if err == nil {
			return nil
		}
		if !IsTemporary(err) {
			return fmt.Errorf("после %d попыток: %w", i+1, err)
		}
		lastErr = err
	}

	return fmt.Errorf("после %d попыток: %w", attempts, lastErr)
}
