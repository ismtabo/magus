package errors

import (
	"fmt"
)

type ValidationError interface {
	error
}

type validationError struct {
	error
}

func NewValidationError(err error) ValidationError {
	return &validationError{err}
}

func (e *validationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.error)
}

func (e *validationError) Unwrap() error {
	return e.error
}
