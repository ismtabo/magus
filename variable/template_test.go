package variable_test

import (
	"testing"

	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/variable"
	"github.com/stretchr/testify/assert"
)

func TestTemplateVariable_Name(t *testing.T) {
	t.Run("it should return the name of the variable", func(t *testing.T) {
		v := variable.NewTemplateVariable("name", "template")

		actual := v.Name()

		assert.Equal(t, "name", actual)
	})
}

func TestTemplateVariable_Value(t *testing.T) {
	t.Run("it should return the value of the template", func(t *testing.T) {
		v := variable.NewTemplateVariable("name", "template")
		ctx := context.New()

		actual, err := v.Value(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "template", actual)
	})

	t.Run("it should return the value of the template with variables", func(t *testing.T) {
		v := variable.NewTemplateVariable("name", "template {{ .var }}")
		vars := immutable.NewMap[string, any](nil).Set("var", "value")
		ctx := context.New().WithVariables(vars)

		actual, err := v.Value(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "template value", actual)
	})

	t.Run("it should return an error if the template is invalid", func(t *testing.T) {
		v := variable.NewTemplateVariable("name", "template {{ .var")
		ctx := context.New()

		_, err := v.Value(ctx)

		assert.Error(t, err)
	})
}
