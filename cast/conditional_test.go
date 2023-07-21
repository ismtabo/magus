package cast_test

import (
	"testing"

	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/source"
	"github.com/ismtabo/magus/template"
	"github.com/ismtabo/magus/variable"
	"github.com/stretchr/testify/assert"
)

func TestConditionalCast_Compile(t *testing.T) {
	t.Run("it should render the source to the destination if the condition is true", func(t *testing.T) {
		src := source.NewSource("Hello World!\n")
		dest := template.NewTemplatedString("testdata/conditional/dest")
		cond := condition.NewAllwaysTrueCondition()
		baseCast := cast.NewBaseCast(src, dest, variable.Variables{})
		c := cast.NewConditionalCast(cond, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewFile("testdata/conditional/dest", "Hello World!\n"),
		}, files)
	})

	t.Run("it should not render the source to the destination if the condition is false", func(t *testing.T) {
		src := source.NewSource("Hello World!\n")
		dest := template.NewTemplatedString("testdata/conditional/dest")
		cond := condition.NewNegatedCondition(condition.NewAllwaysTrueCondition())
		baseCast := cast.NewBaseCast(src, dest, variable.Variables{})
		c := cast.NewConditionalCast(cond, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{}, files)
	})
}
