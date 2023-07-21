package errors

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
	return "render error"
}

func (e *renderError) Unwrap() error {
	return e.error
}
