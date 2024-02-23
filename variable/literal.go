package variable

import "github.com/ismtabo/magus/v2/context"

var _ Variable = &LiteralVariable{}

// LiteralVariable is a variable that has a literal value.
type LiteralVariable struct {
	name  string
	value any
}

// NewLiteralVariable creates a new literal variable.
func NewLiteralVariable(name string, value any) Variable {
	return &LiteralVariable{
		name:  name,
		value: value,
	}
}

// Name returns the name of the variable.
func (v *LiteralVariable) Name() string {
	return v.name
}

// Value returns the value of the variable.
func (v *LiteralVariable) Value(ctx context.Context) (any, error) {
	return v.value, nil
}
