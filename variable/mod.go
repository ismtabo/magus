package variable

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/manifest"
)

// Variable is a variable that can be used in a template.
type Variable interface {
	// Name returns the name of the variable.
	Name() string
	// Value returns the value of the variable.
	Value(ctx context.Context) (any, error)
}

// New creates a new variable.
func NewVariable(variable manifest.Variable) Variable {
	name := variable.Name
	if variable.Env != "" {
		return NewEnvironmentVariable(name, variable.Env)
	}
	if variable.Template != "" {
		return NewTemplateVariable(name, variable.Template)
	}
	return NewLiteralVariable(name, variable.Value)
}

type Variables = []Variable

// NewVariables creates a new slice of variables.
func NewVariables(variables []manifest.Variable) Variables {
	vars := make([]Variable, len(variables))
	for i, variable := range variables {
		vars[i] = NewVariable(variable)
	}
	return vars
}
