package myErrors

import "errors"

func Chain(err error) []string {
	if err == nil {
		return []string{}
	}

	var result []string
	current := err

	for current != nil {
		result = append(result, current.Error())
		current = errors.Unwrap(current)
	}

	return result
}
