package cast

import (
	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
)

var _ Cast = (*ConditionalCast)(nil)

// ConditionalCast is a cast that renders files from a source to a destination if a condition is met.
type ConditionalCast struct {
	baseCast *BaseCast
	cond     condition.Condition
}

// NewConditionalCast creates a new conditional cast.
func NewConditionalCast(condition condition.Condition, baseCast *BaseCast) *ConditionalCast {
	return &ConditionalCast{
		baseCast: baseCast,
		cond:     condition,
	}
}

// Compile compiles the cast.
func (c *ConditionalCast) Compile(ctx context.Context) ([]file.File, error) {
	val, err := c.cond.Evaluate(ctx)
	if err != nil {
		return nil, err
	}
	if !val {
		return []file.File{}, nil
	}
	return c.baseCast.Compile(ctx)
}
