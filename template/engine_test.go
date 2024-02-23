package template_test

import (
	"testing"

	go_errors "errors"

	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/template"
	"github.com/stretchr/testify/assert"
)

func TestTemplateEngine_Render(t *testing.T) {
	t.Run("it should return the render of the template", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		vars = vars.Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		tmpl := "Hello {{ .name }}"
		expected := "Hello John"

		actual, err := template.Engine.Render(ctx, tmpl)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("it should return error if template is invalid", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		tmpl := "Hello {{ .name }"

		_, actual := template.Engine.Render(ctx, tmpl)

		assert.Error(t, actual)
	})

	t.Run("it should return error if the render fails", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		helpers := immutable.NewMap[string, any](nil)
		helpers = helpers.Set("throw", func() (string, error) { return "", go_errors.New("error") })
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		ctx = ctx.WithHelpers(helpers)

		tmpl := "Hello {{ throw }}"

		result, actual := template.Engine.Render(ctx, tmpl)

		t.Log(result)
		assert.Error(t, actual)
	})
}
