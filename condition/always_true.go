package condition

import "github.com/ismtabo/magus/v2/context"

// AlwaysTrueCondition is a condition that allways evaluates to true.
type AlwaysTrueCondition struct{}

// NewAlwaysTrueCondition creates a new condition that allways evaluates to true.
func NewAlwaysTrueCondition() *AlwaysTrueCondition {
	return &AlwaysTrueCondition{}
}

// Evaluate evaluates the condition.
func (c *AlwaysTrueCondition) Evaluate(ctx context.Context) (bool, error) {
	return true, nil
}
