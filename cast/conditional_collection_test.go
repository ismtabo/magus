package cast_test

import (
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

func TestConditionalCollectionCast_Compile(t *testing.T) {
	t.Run("it should return an empty list of files if the condition is false", func(t *testing.T) {
		src := source.NewTemplateSource("Hello World!")
		dest := template.NewTemplatedPath("path/to/dest")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		cond := condition.NewNegatedCondition(condition.NewAlwaysTrueCondition())
		condCast := cast.NewConditionalCast(cond, baseCast)
		coll := template.NewTemplatedString("[1, 2, 3]")
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		collCast := cast.NewCollectionCast(coll, alias, filter, baseCast)
		c := cast.NewConditionalCollectionCast(condCast, collCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{}, files)
	})

	t.Run("it should return the collection if the condition is true", func(t *testing.T) {
		src := source.NewTemplateSource("Hello {{ .It }}!")
		dest := template.NewTemplatedPath("path/to/dest/{{ .It }}")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		cond := condition.NewAlwaysTrueCondition()
		condCast := cast.NewConditionalCast(cond, baseCast)
		coll := template.NewTemplatedString(`["harry", "ron", "hermione"]`)
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		collCast := cast.NewCollectionCast(coll, alias, filter, baseCast)
		c := cast.NewConditionalCollectionCast(condCast, collCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("path/to/dest/harry", "Hello harry!"),
			file.NewTextFile("path/to/dest/ron", "Hello ron!"),
			file.NewTextFile("path/to/dest/hermione", "Hello hermione!"),
		}, files)
	})

	t.Run("it should return an error if the condition cast returns an error", func(t *testing.T) {
		src := source.NewTemplateSource("Hello World!")
		dest := template.NewTemplatedPath("path/to/dest")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		cond := NewFailingCondition()
		condCast := cast.NewConditionalCast(cond, baseCast)
		coll := template.NewTemplatedString("[1, 2, 3]")
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		collCast := cast.NewCollectionCast(coll, alias, filter, baseCast)
		c := cast.NewConditionalCollectionCast(condCast, collCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.Error(t, err)
		assert.Nil(t, files)
	})
}
