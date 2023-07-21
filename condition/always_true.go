package condition

import "github.com/ismtabo/magus/context"

// AllwaysTrueCondition is a condition that allways evaluates to true.
type AllwaysTrueCondition struct{}

// NewAllwaysTrueCondition creates a new condition that allways evaluates to true.
func NewAllwaysTrueCondition() *AllwaysTrueCondition {
	return &AllwaysTrueCondition{}
}

// Evaluate evaluates the condition.
func (c *AllwaysTrueCondition) Evaluate(ctx context.Context) (bool, error) {
	return true, nil
}
