package cast_test

import (
	"errors"
	"testing"

	"github.com/ismtabo/magus/v2/cast"
	"github.com/ismtabo/magus/v2/condition"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/source"
	"github.com/ismtabo/magus/v2/template"
	"github.com/ismtabo/magus/v2/variable"
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
		src := source.NewTemplateSource("Hello World!\n")
		dest := template.NewTemplatedPath("testdata/conditional/dest")
		cond := condition.NewAlwaysTrueCondition()
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
		src := source.NewTemplateSource("Hello World!\n")
		dest := template.NewTemplatedPath("testdata/conditional/dest")
		cond := condition.NewNegatedCondition(condition.NewAlwaysTrueCondition())
		baseCast := cast.NewBaseCast(src, dest, variable.Variables{})
		c := cast.NewConditionalCast(cond, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{}, files)
	})

	t.Run("it should return an error if the condition evaluation fails", func(t *testing.T) {
		src := source.NewTemplateSource("Hello World!\n")
		dest := template.NewTemplatedPath("testdata/conditional/dest")
		cond := NewFailingCondition()
		baseCast := cast.NewBaseCast(src, dest, variable.Variables{})
		c := cast.NewConditionalCast(cond, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.Error(t, err)
		assert.Nil(t, files)
	})
}
