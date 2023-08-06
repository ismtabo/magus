package condition_test

import (
	"testing"

	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/stretchr/testify/assert"
)

func TestTemplatedCondition_Evaluate(t *testing.T) {
	t.Run("it should return true if the condition is empty", func(t *testing.T) {
		c := condition.NewTemplateCondition("")
		ctx := context.New()

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("it should return the result of the condition", func(t *testing.T) {
		c := condition.NewTemplateCondition("{{ .name | len | ne 0 }}")
		vars := immutable.NewMap[string, any](nil).Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithVariables(vars)

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("it should return true if the condition is true", func(t *testing.T) {
		c := condition.NewTemplateCondition("true")
		vars := immutable.NewMap[string, any](nil).Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithVariables(vars)

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("it should return false if the condition is not true", func(t *testing.T) {
		c := condition.NewTemplateCondition("false")
		ctx := context.New()

		result, err := c.Evaluate(ctx)

		assert.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("it should return an error if the condition is invalid", func(t *testing.T) {
		c := condition.NewTemplateCondition("{{ .name | len > 0}")
		ctx := context.New()

		_, err := c.Evaluate(ctx)

		assert.Error(t, err)
	})
}
