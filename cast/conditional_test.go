package cast_test

import (
	"errors"
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

type FailingCondition struct{}

func (c *FailingCondition) Evaluate(ctx context.Context) (bool, error) {
	return false, errors.New("condition evaluation failed")
}

func NewFailingCondition() condition.Condition {
	return &FailingCondition{}
}

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
			file.NewTextFile("testdata/conditional/dest", "Hello World!\n"),
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

	t.Run("it should return an error if the condition evaluation fails", func(t *testing.T) {
		src := source.NewSource("Hello World!\n")
		dest := template.NewTemplatedString("testdata/conditional/dest")
		cond := NewFailingCondition()
		baseCast := cast.NewBaseCast(src, dest, variable.Variables{})
		c := cast.NewConditionalCast(cond, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.Error(t, err)
		assert.Nil(t, files)
	})
}
