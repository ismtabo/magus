package filetree_test

import (
	"testing"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/filetree"
	"github.com/stretchr/testify/assert"
)

func TestAssertNotHaveWriteConficts(t *testing.T) {
	t.Run("should not return error if no write conflicts", func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithCwd(t.TempDir())
		files := []file.File{
			file.NewTextFile("foo/bar/baz", "foo"),
			file.NewTextFile("foo/bar/qux", "foo"),
			file.NewTextFile("foo/baz", "foo"),
			file.NewTextFile("foo/qux", "foo"),
			file.NewTextFile("bar/baz", "foo"),
			file.NewTextFile("bar/qux", "foo"),
			file.NewTextFile("baz", "foo"),
			file.NewTextFile("qux", "foo"),
		}
		assert.NoError(t, filetree.AssertNotHaveWriteConficts(ctx, files))
	})

	t.Run("should return error if write conflict", func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithCwd(t.TempDir())
		files := []file.File{
			file.NewTextFile("foo/bar/baz", "foo"),
			file.NewTextFile("foo/bar", "foo"),
		}
		assert.Error(t, filetree.AssertNotHaveWriteConficts(ctx, files))
	})
}
