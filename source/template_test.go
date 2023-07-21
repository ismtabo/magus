package source_test

import (
	"testing"

	go_errors "errors"

	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/source"
	"github.com/stretchr/testify/assert"
)

func TestTemplateSource_Compile(t *testing.T) {
	t.Run("when the template is valid", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		vars = vars.Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithCwd("/tmp")
		ctx = ctx.WithVariables(vars)
		tmpl := "Hello {{ .name }}"
		expected := []file.File{{
			Path:  "/tmp",
			Value: "Hello John",
		}}

		actual, err := source.NewTemplateSource(tmpl).Compile(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("when the template is invalid", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		vars = vars.Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		tmpl := "Hello {{ .name }"

		_, actual := source.NewTemplateSource(tmpl).Compile(ctx)

		assert.EqualError(t, actual, "validation error")
	})

	t.Run("when render fails", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		helpers := immutable.NewMap[string, any](nil)
		helpers = helpers.Set("throw", func() (string, error) { return "", go_errors.New("error") })
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		ctx = ctx.WithHelpers(helpers)

		tmpl := "Hello {{ throw }}"

		result, actual := source.NewTemplateSource(tmpl).Compile(ctx)

		t.Log(result)
		assert.EqualError(t, actual, "render error")
	})
}
