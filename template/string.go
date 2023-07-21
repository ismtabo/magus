package template

import "github.com/ismtabo/magus/context"

// TemplatedString is a string that can be templated.
type TemplatedString struct {
	string
}

// NewTemplatedString creates a new TemplatedString.
func NewTemplatedString(str string) TemplatedString {
	return TemplatedString{str}
}

func (s TemplatedString) Compile(ctx context.Context) (string, error) {
	return Engine.Render(ctx, s.string)
}
