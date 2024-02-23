package template_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/template"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestTemplatePath_Render(t *testing.T) {
	t.Run("it should return the rendered path", func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithVariable("foo", "bar")
		path := "foo/{{ .foo }}/baz"
		expected := "foo/bar/baz"
		tmpl := template.NewTemplatedPath(path)

		actual, err := tmpl.Render(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("it should return an error if the template is invalid", func(t *testing.T) {
		ctx := context.New()
		path := "foo/{{ .foo }/baz"
		tmpl := template.NewTemplatedPath(path)

		_, err := tmpl.Render(ctx)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the rendered path is not local", func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithVariable("foo", "bar")
		path := "/{{ .foo }}"
		tmpl := template.NewTemplatedPath(path)

		_, err := tmpl.Render(ctx)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the template is not a path", func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithVariable("foo", "bar")
		path := dedent.Dedent(`
		# Hello World
		"{{ .foo }}"
		`)
		tmpl := template.NewTemplatedPath(path)

		_, err := tmpl.Render(ctx)

		assert.Error(t, err)
	})
}
