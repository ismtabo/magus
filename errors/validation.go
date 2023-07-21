package errors

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
	return "validation error"
}

func (e *validationError) Unwrap() error {
	return e.error
}
