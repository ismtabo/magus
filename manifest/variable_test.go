package manifest_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/manifest"
	"github.com/ismtabo/magus/v2/variable"
	"github.com/stretchr/testify/assert"
)

func TestNewVariable(t *testing.T) {
	t.Run("it should return a new literal variable", func(t *testing.T) {
		v := manifest.Variable{
			Name:  "name",
			Value: "value",
		}
		actual := manifest.NewVariable(v)

		assert.Implements(t, (*variable.Variable)(nil), actual)
		assert.IsType(t, (*variable.LiteralVariable)(nil), actual)
		assert.Equal(t, "name", actual.Name())
	})

	t.Run("it should return a new environment variable", func(t *testing.T) {
		v := manifest.Variable{
			Name: "name",
			Env:  "env",
		}
		actual := manifest.NewVariable(v)

		assert.Implements(t, (*variable.Variable)(nil), actual)
		assert.IsType(t, (*variable.EnvironmentVariable)(nil), actual)
		assert.Equal(t, "name", actual.Name())
	})

	t.Run("it should return a new template variable", func(t *testing.T) {
		v := manifest.Variable{
			Name:     "name",
			Template: "template",
		}
		actual := manifest.NewVariable(v)

		assert.Implements(t, (*variable.Variable)(nil), actual)
		assert.IsType(t, (*variable.TemplateVariable)(nil), actual)
		assert.Equal(t, "name", actual.Name())
	})

	t.Run("it should return a new literal variable if no type is specified", func(t *testing.T) {
		v := manifest.Variable{
			Name: "name",
		}
		actual := manifest.NewVariable(v)

		assert.Implements(t, (*variable.Variable)(nil), actual)
		assert.IsType(t, (*variable.LiteralVariable)(nil), actual)
		assert.Equal(t, "name", actual.Name())
	})
}

func TestNewVariables(t *testing.T) {
	t.Run("it should return a new list of variables", func(t *testing.T) {
		vs := []manifest.Variable{
			{
				Name:  "name1",
				Value: "value1",
			},
			{
				Name:  "name2",
				Value: "value2",
			},
		}
		actual := manifest.NewVariables(vs)

		assert.Len(t, actual, 2)
		assert.Equal(t, "name1", actual[0].Name())
		assert.Equal(t, "name2", actual[1].Name())
	})
}
