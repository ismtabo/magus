package variable

import (
	"os"

	"github.com/ismtabo/magus/v2/context"
)

var _ Variable = &EnvironmentVariable{}

// EnvironmentVariable is a variable that has a value from an environment variable.
type EnvironmentVariable struct {
	name string
	env  string
}

// NewEnvironmentVariable creates a new environment variable.
func NewEnvironmentVariable(name string, env string) Variable {
	return &EnvironmentVariable{
		name: name,
		env:  env,
	}
}

// Name returns the name of the variable.
func (v *EnvironmentVariable) Name() string {
	return v.name
}

// Value returns the value of the variable.
func (v *EnvironmentVariable) Value(ctx context.Context) (any, error) {
	return os.Getenv(v.env), nil
}
