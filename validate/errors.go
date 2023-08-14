package validate

import "github.com/pkg/errors"

type ValidationError = error

const (
	ErrInvalidVersion = "invalid version"
)

func WrapInvalidVersionErrorf(message string, args ...interface{}) ValidationError {
	err := errors.Errorf(message, args...)
	err = errors.Wrap(err, ErrInvalidVersion)
	return err
}
