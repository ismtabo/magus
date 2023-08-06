package condition

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/template"
)

// TemplateCondition is a condition that evaluates a template.
type TemplateCondition struct {
	string
}

// NewTemplateCondition creates a new condition that evaluates a template.
func NewTemplateCondition(condition string) *TemplateCondition {
	return &TemplateCondition{condition}
}

// Evaluate evaluates the condition.
func (c *TemplateCondition) Evaluate(ctx context.Context) (bool, error) {
	if c.string == "" {
		return true, nil
	}
	val, err := template.Engine.Render(ctx, c.string)
	if err != nil {
		return false, err
	}
	return val == "true", nil
}
