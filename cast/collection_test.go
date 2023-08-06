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

func TestCollection_Compile(t *testing.T) {
	t.Run("it should return an error if the collection fails to compile", func(t *testing.T) {
		src := source.NewTemplateSource("Hello World!")
		dest := template.NewTemplatedPath("path/to/dest")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		coll := template.NewTemplatedString("{{ .It }")
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		c := cast.NewCollectionCast(coll, alias, filter, baseCast)
		ctx := context.New()

		_, err := c.Compile(ctx)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the collection is not a valid JSON array", func(t *testing.T) {
		src := source.NewTemplateSource("Hello World!")
		dest := template.NewTemplatedPath("path/to/dest")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		coll := template.NewTemplatedString("Hello World!")
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		c := cast.NewCollectionCast(coll, alias, filter, baseCast)
		ctx := context.New()

		_, err := c.Compile(ctx)

		assert.Error(t, err)
	})

	t.Run("it should return the rendered files for the collection", func(t *testing.T) {
		src := source.NewTemplateSource("Hello {{ .It }}!")
		dest := template.NewTemplatedPath("path/to/dest/{{ .It }}")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		coll := template.NewTemplatedString(`["harry", "ron", "hermione"]`)
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		c := cast.NewCollectionCast(coll, alias, filter, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("path/to/dest/harry", "Hello harry!"),
			file.NewTextFile("path/to/dest/ron", "Hello ron!"),
			file.NewTextFile("path/to/dest/hermione", "Hello hermione!"),
		}, files)
	})

	t.Run("it should return an error if the filter fails to evaluate", func(t *testing.T) {
		src := source.NewTemplateSource("Hello World!")
		dest := template.NewTemplatedPath("path/to/dest")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		coll := template.NewTemplatedString(`["harry", "ron", "hermione"]`)
		alias := "It"
		filter := NewFailingCondition()
		c := cast.NewCollectionCast(coll, alias, filter, baseCast)
		ctx := context.New()

		_, err := c.Compile(ctx)

		assert.Error(t, err)
	})

	t.Run("it should return the rendered files for the collection if the filter is true", func(t *testing.T) {
		src := source.NewTemplateSource("Hello {{ .It }}!")
		dest := template.NewTemplatedPath("path/to/dest/{{ .It }}")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		coll := template.NewTemplatedString(`["harry", "ron", "hermione"]`)
		alias := "It"
		filter := condition.NewAlwaysTrueCondition()
		c := cast.NewCollectionCast(coll, alias, filter, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{
			file.NewTextFile("path/to/dest/harry", "Hello harry!"),
			file.NewTextFile("path/to/dest/ron", "Hello ron!"),
			file.NewTextFile("path/to/dest/hermione", "Hello hermione!"),
		}, files)
	})

	t.Run("it should return an empty list of files if the filter is false", func(t *testing.T) {
		src := source.NewTemplateSource("Hello {{ .It }}!")
		dest := template.NewTemplatedPath("path/to/dest/{{ .It }}")
		vars := variable.Variables{}
		baseCast := cast.NewBaseCast(src, dest, vars)
		coll := template.NewTemplatedString(`["harry", "ron", "hermione"]`)
		alias := "It"
		filter := condition.NewNegatedCondition(condition.NewAlwaysTrueCondition())
		c := cast.NewCollectionCast(coll, alias, filter, baseCast)
		ctx := context.New()

		files, err := c.Compile(ctx)

		assert.NoError(t, err)
		assert.Equal(t, []file.File{}, files)
	})
}
