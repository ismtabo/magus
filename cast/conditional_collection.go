package cast

import (
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
)

var _ Cast = (*ConditionalCollectionCast)(nil)

// ConditionalCollectionCast is a cast that renders files from a source to a destination for a given collection if a condition is met.
type ConditionalCollectionCast struct {
	*CollectionCast
	*ConditionalCast
}

// NewConditionalCollectionCast creates a new conditional collection cast.
func NewConditionalCollectionCast(condCast *ConditionalCast, collCast *CollectionCast) *ConditionalCollectionCast {
	return &ConditionalCollectionCast{
		CollectionCast:  collCast,
		ConditionalCast: condCast,
	}
}

// Compile compiles the cast.
func (c *ConditionalCollectionCast) Compile(ctx context.Context) ([]file.File, error) {
	val, err := c.cond.Evaluate(ctx)
	if err != nil {
		return nil, err
	}
	if !val {
		return []file.File{}, nil
	}
	return c.CollectionCast.Compile(ctx)
}
