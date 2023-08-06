package condition_test

import (
	"errors"
	"testing"

	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/stretchr/testify/assert"
)

type FailingCondition struct{}

func (c *FailingCondition) Evaluate(ctx context.Context) (bool, error) {
	return false, errors.New("condition evaluation failed")
}

func NewFailingCondition() condition.Condition {
	return &FailingCondition{}
}

func TestNegatedCondition_Evaluate(t *testing.T) {
	t.Run("it should return the negated result of the condition", func(t *testing.T) {
		c := condition.NewNegatedCondition(condition.NewAlwaysTrueCondition())
		ctx := context.New()

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("it should return an error if the condition evaluation fails", func(t *testing.T) {
		c := condition.NewNegatedCondition(NewFailingCondition())
		ctx := context.New()

		result, err := c.Evaluate(ctx)

		assert.Error(t, err)
		assert.False(t, result)
	})
}
