package myErrors

import (
	"errors"
	"strings"
)

var (
	errBadName  = errors.New("Bad Name")
	errBadEmail = errors.New("Bad Email")
	errBadAge   = errors.New("Bad Age")
)

func ValidateForm(name, email string, age int) error {

	var errs []error

	if name == "" {
		errs = append(errs, errBadName)
	} else {
		for _, letter := range name {
			if (letter < 'a' || letter > 'z') && (letter < 'A' || letter > 'Z') && letter != ' ' && letter != '-' {
				errs = append(errs, errBadName)
				break
			}
		}
	}

	if email == "" {
		errs = append(errs, errBadEmail)
	} else {
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			errs = append(errs, errBadEmail)

		}
	}

	if age < 0 || age > 120 {
		errs = append(errs, errBadAge)
	}

	return errors.Join(errs...)

}
