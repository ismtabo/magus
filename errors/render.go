package errors

import "fmt"

type RenderError interface {
	error
}

type renderError struct {
	error
}

func NewRenderError(err error) RenderError {
	return &renderError{err}
}

func (e *renderError) Error() string {
	return fmt.Sprintf("render error: %s", e.error)
}

func (e *renderError) Unwrap() error {
	return e.error
}
