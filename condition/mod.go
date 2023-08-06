package condition

import "github.com/ismtabo/magus/context"

// Condition is a condition that can be evaluated.
type Condition interface {
	// Evaluate evaluates the condition.
	Evaluate(ctx context.Context) (bool, error)
}
