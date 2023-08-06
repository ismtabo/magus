package template_test

import (
	"testing"

	go_errors "errors"

	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/template"
	"github.com/stretchr/testify/assert"
)

func TestTemplatedString_Render(t *testing.T) {
	t.Run("it should return the render of the template", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		vars = vars.Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		tmpl := template.NewTemplatedString("Hello {{ .name }}")
		expected := "Hello John"

		actual, err := tmpl.Render(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("it should return error if template is invalid", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		tmpl := template.NewTemplatedString("Hello {{ .name }")

		_, actual := tmpl.Render(ctx)

		assert.Error(t, actual)
	})

	t.Run("it should return error if the render fails", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		helpers := immutable.NewMap[string, any](nil)
		helpers = helpers.Set("throw", func() (string, error) { return "", go_errors.New("error") })
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		ctx = ctx.WithHelpers(helpers)

		tmpl := template.NewTemplatedString("Hello {{ throw }}")

		result, actual := tmpl.Render(ctx)

		t.Log(result)
		assert.Error(t, actual)
	})
}
