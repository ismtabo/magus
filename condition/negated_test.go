package condition_test

import (
	"testing"

	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/stretchr/testify/assert"
)

func TestNegatedCondition_Evaluate(t *testing.T) {
	t.Run("it should return the negated result of the condition", func(t *testing.T) {
		c := condition.NewNegatedCondition(condition.NewAllwaysTrueCondition())
		ctx := context.New()

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.False(t, result)
	})
}
