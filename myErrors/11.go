package myErrors

import "fmt"

type OutOfRange struct {
	Value, Min, Max int
}

func (e *OutOfRange) Error() string {

	return fmt.Sprintf("%d out of range: [%d, %d]", e.Value, e.Min, e.Max)
}
