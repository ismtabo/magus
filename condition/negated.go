package condition

import "github.com/ismtabo/magus/v2/context"

// NegatedCondition is a condition that negates another condition.
type NegatedCondition struct {
	Condition
}

// NewNegatedCondition creates a new condition that negates the given condition.
func NewNegatedCondition(cond Condition) *NegatedCondition {
	return &NegatedCondition{cond}
}

// Evaluate evaluates the condition.
func (c *NegatedCondition) Evaluate(ctx context.Context) (bool, error) {
	val, err := c.Condition.Evaluate(ctx)
	if err != nil {
		return false, err
	}
	return !val, nil
}
