package main

import (
	"errors"
	"fmt"
	"strconv"
)

type myError interface {
	Error() string
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func ReadAge(s string) (int, error) {
	c, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %s", s)
	}

	if c < 0 || c > 150 {
		return 0, &OutOfRange{
			Value: c,
			Min:   0,
			Max:   150,
		}
	}
	return c, nil
}

func process(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "err"
	}
	if n <= 0 {
		return "neg"
	}
	if n >= 100 {
		return "big"
	}

	return "ok"
}

type OutOfRange struct {
	Value, Min, Max int
}

func (e *OutOfRange) Error() string {
	return fmt.Sprintf("%d out of range: [%d, %d]", e.Value, e.Min, e.Max)
}

var ErrNotFound = errors.New("не найдено")

func Find(m map[string]int, k string) (int, error) {
	val, ok := m[k]
	if !ok {
		return 0, ErrNotFound
	}
	return val, nil
}
