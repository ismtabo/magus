package magic_test

import (
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/v2/cast"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/magic"
	"github.com/ismtabo/magus/v2/source"
	"github.com/ismtabo/magus/v2/template"
	"github.com/ismtabo/magus/v2/variable"
	"github.com/stretchr/testify/assert"
)

func TestMagic_Render(t *testing.T) {
	t.Run("should render a magic", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{}, []cast.Cast{})

		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, files, []file.File{})
	})

	t.Run("should render a magic with variables", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{
			variable.NewLiteralVariable("name", "John Doe"),
		}, []cast.Cast{})

		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, files, []file.File{})
	})

	t.Run("should render a magic with casts", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("test"), template.NewTemplatedPath("test"), variable.Variables{}),
		})

		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("test", "test"),
		}, files)
	})

	t.Run("should render a magic with variables and casts", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{
			variable.NewLiteralVariable("name", "John Doe"),
		}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("{{ .name }}"), template.NewTemplatedPath("test"), variable.Variables{}),
		})
		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("test", "John Doe"),
		}, files)
	})

	t.Run("should render a magic with variables in options", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("{{ .name }}"), template.NewTemplatedPath("test"), variable.Variables{}),
		})
		files, err := m.Render(context.New(), magic.MagicRenderOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("name", "John Doe"),
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("test", "John Doe"),
		}, files)
	})

	t.Run("should return an error if a variable fails to render", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{
			variable.NewTemplateVariable("age", "{{ .name }"),
		}, []cast.Cast{})

		_, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.Error(t, err)
	})

	t.Run("should return an error if a cast fails to render", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("{{ .name }"), template.NewTemplatedPath("test"), variable.Variables{}),
		})

		_, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.Error(t, err)
	})
}

func TestMagic_Compile(t *testing.T) {
	t.Run("should compile a magic", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{}, []cast.Cast{})

		_, err := m.Compile(context.New(), "")

		assert.NoError(t, err)
	})

	t.Run("should render a magic with dest path", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("test"), template.NewTemplatedPath("test"), variable.Variables{}),
		})

		files, err := m.Compile(context.New(), "out")

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile(filepath.Join("out", "test"), "test"),
		}, files)
	})

	t.Run("should return an error if a variable fails to render", func(t *testing.T) {
		m := magic.NewMagic("test", []variable.Variable{
			variable.NewTemplateVariable("age", "{{ .name }"),
		}, []cast.Cast{})

		_, err := m.Compile(context.New(), "")

		assert.Error(t, err)
	})
}
