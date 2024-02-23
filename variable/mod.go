package variable

import (
	"github.com/ismtabo/magus/v2/context"
)

// Variable is a variable that can be used in a template.
type Variable interface {
	// Name returns the name of the variable.
	Name() string
	// Value returns the value of the variable.
	Value(ctx context.Context) (any, error)
}

type Variables = []Variable
