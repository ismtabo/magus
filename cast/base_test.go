package cast_test

import (
	"testing"

	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/source"
	"github.com/ismtabo/magus/template"
	"github.com/ismtabo/magus/variable"
	"github.com/stretchr/testify/assert"
)

type FailingVariable struct{}

func (v *FailingVariable) Name() string {
	return "failing"
}

func (v *FailingVariable) Value(ctx context.Context) (interface{}, error) {
	return nil, assert.AnError
}

func TestBaseCast_Compile(t *testing.T) {
	t.Run("it should render the source to the destination", func(t *testing.T) {
		src := source.NewSource("Hello {{ .name }}!\n")
		dest := template.NewTemplatedString("testdata/base/dest")
		vars := []variable.Variable{
			variable.NewLiteralVariable("name", "John"),
		}
		c := cast.NewBaseCast(src, dest, vars)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewFile("testdata/base/dest", "Hello John!\n"),
		}, files)
	})

	t.Run("it should render the source to the destination with the correct cwd", func(t *testing.T) {
		src := source.NewSource("Hello {{ .name }}!\n")
		dest := template.NewTemplatedString("./testdata/base/dest")
		vars := []variable.Variable{
			variable.NewLiteralVariable("name", "John"),
		}
		c := cast.NewBaseCast(src, dest, vars)
		ctx := context.New()
		ctx = ctx.WithCwd("root")

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewFile("root/testdata/base/dest", "Hello John!\n"),
		}, files)
	})

	t.Run("it should render the source to the destination with correct variables", func(t *testing.T) {
		src := source.NewSource("Hello {{ .name }}!\n")
		dest := template.NewTemplatedString("./testdata/base/dest/{{ .filename }}")
		vars := []variable.Variable{
			variable.NewLiteralVariable("name", "John"),
			variable.NewLiteralVariable("filename", "john"),
		}
		c := cast.NewBaseCast(src, dest, vars)
		ctx := context.New()
		ctx = ctx.WithCwd("root")

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewFile("root/testdata/base/dest/john", "Hello John!\n"),
		}, files)
	})

	t.Run("should return error if variables evaluation fails", func(t *testing.T) {
		src := source.NewSource("Hello {{ .name }}!\n")
		dest := template.NewTemplatedString("./testdata/base/dest/{{ .filename }}")
		vars := []variable.Variable{
			variable.NewLiteralVariable("name", "John"),
			&FailingVariable{},
		}
		c := cast.NewBaseCast(src, dest, vars)
		ctx := context.New()
		ctx = ctx.WithCwd("root")

		files, err := c.Compile(ctx)

		assert.Error(t, err)
		assert.Nil(t, files)
	})

	t.Run("should return error if destination evaluation fails", func(t *testing.T) {
		src := source.NewSource("Hello {{ .name }}!\n")
		dest := template.NewTemplatedString("./testdata/base/dest/{{ .filename }")
		vars := []variable.Variable{
			variable.NewLiteralVariable("name", "John"),
			variable.NewLiteralVariable("filename", "john"),
		}
		c := cast.NewBaseCast(src, dest, vars)
		ctx := context.New()
		ctx = ctx.WithCwd("root")
		ctx = ctx.WithVariables(ctx.Variables().Set("filename", assert.AnError))

		files, err := c.Compile(ctx)

		assert.Error(t, err)
		assert.Nil(t, files)
	})
}
