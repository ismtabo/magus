package source_test

import (
	"testing"

	go_errors "errors"

	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/source"
	"github.com/stretchr/testify/assert"
)

func TestTemplateSource_Compile(t *testing.T) {
	t.Run("it should return compiled file", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		vars = vars.Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithCwd("/tmp")
		ctx = ctx.WithVariables(vars)
		tmpl := "Hello {{ .name }}"

		actual, err := source.NewTemplateSource(tmpl).Compile(ctx, "/tmp/file.txt")

		assert.Nil(t, err)
		assert.Equal(t, []file.File{file.NewTextFile("/tmp/file.txt", "Hello John")}, actual)
	})

	t.Run("it should return error if template is invalid", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		vars = vars.Set("name", "John")
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		tmpl := "Hello {{ .name }"

		_, actual := source.NewTemplateSource(tmpl).Compile(ctx, "/tmp/file.txt")

		assert.Error(t, actual)
	})

	t.Run("it should return error if render fails", func(t *testing.T) {
		vars := immutable.NewMap[string, any](nil)
		helpers := immutable.NewMap[string, any](nil)
		helpers = helpers.Set("throw", func() (string, error) { return "", go_errors.New("error") })
		ctx := context.New()
		ctx = ctx.WithVariables(vars)
		ctx = ctx.WithHelpers(helpers)

		tmpl := "Hello {{ throw }}"

		result, actual := source.NewTemplateSource(tmpl).Compile(ctx, "/tmp/file.txt")

		t.Log(result)
		assert.Error(t, actual)
	})
}
