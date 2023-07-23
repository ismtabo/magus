package variable

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/template"
)

var _ Variable = &TemplateVariable{}

// TemplateVariable is a variable that has a value from a template.
type TemplateVariable struct {
	name string
	tmpl template.TemplatedString
}

// NewTemplateVariable creates a new template variable.
func NewTemplateVariable(name string, tmpl string) Variable {
	return &TemplateVariable{
		name: name,
		tmpl: template.NewTemplatedString(tmpl),
	}
}

// Name returns the name of the variable.
func (v *TemplateVariable) Name() string {
	return v.name
}

// Value returns the value of the variable.
func (v *TemplateVariable) Value(ctx context.Context) (any, error) {
	return v.tmpl.Render(ctx)
}
