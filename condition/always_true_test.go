package condition_test

import (
	"testing"

	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrueCondition_Evaluate(t *testing.T) {
	t.Run("it should always return true", func(t *testing.T) {
		c := condition.NewAlwaysTrueCondition()
		ctx := context.New()

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.True(t, result)
	})
}
