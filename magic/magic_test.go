package magic_test

import (
	"testing"

	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/source"
	"github.com/ismtabo/magus/template"
	"github.com/ismtabo/magus/variable"
	"github.com/stretchr/testify/assert"
)

func TestMagic_Render(t *testing.T) {
	t.Run("should render a magic", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{}, []cast.Cast{})

		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, files, []file.File{})
	})

	t.Run("should render a magic with variables", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{
			variable.NewLiteralVariable("name", "John Doe"),
		}, []cast.Cast{})

		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, files, []file.File{})
	})

	t.Run("should render a magic with casts", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("test"), template.NewTemplatedPath("test"), variable.Variables{}),
		})

		files, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("test", "test"),
		}, files)
	})

	t.Run("should render a magic with variables and casts", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{
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

	t.Run("should render a magic with variables and casts and render options", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("{{ .name }}"), template.NewTemplatedPath("test"), variable.Variables{}),
		})
		files, err := m.Render(context.New(), magic.MagicRenderOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("name", "Jane Doe"),
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("test", "Jane Doe"),
		}, files)
	})

	t.Run("should render option variables before magic variables", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{
			variable.NewTemplateVariable("name", "{{ .name }}"),
		}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("{{ .name }}"), template.NewTemplatedPath("test"), variable.Variables{}),
		})
		files, err := m.Render(context.New(), magic.MagicRenderOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("name", "Jane Doe"),
			},
		})

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("test", "Jane Doe"),
		}, files)
	})

	t.Run("should return an error if a variable fails to render", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{
			variable.NewTemplateVariable("age", "{{ .name }"),
		}, []cast.Cast{})

		_, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.Error(t, err)
	})

	t.Run("should return an error if a cast fails to render", func(t *testing.T) {
		m := magic.NewMagic("1.0.0", "test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("{{ .name }"), template.NewTemplatedPath("test"), variable.Variables{}),
		})

		_, err := m.Render(context.New(), magic.MagicRenderOptions{})

		assert.Error(t, err)
	})
}
