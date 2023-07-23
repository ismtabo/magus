package file_test

import (
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/stretchr/testify/assert"
)

func TestFile_Abs(t *testing.T) {
	t.Run("it should return the absolute path of the file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "path")
		f := file.NewFile(path, []byte{})
		ctx := context.New()

		abs, err := f.Abs(ctx)

		assert.NoError(t, err)
		assert.Equal(t, path, abs)
	})

	t.Run("it should return absolute path of the file with context cwd", func(t *testing.T) {
		f := file.NewFile("path", []byte{})
		dir := t.TempDir()
		ctx := context.New()
		ctx = ctx.WithCwd(dir)

		abs, err := f.Abs(ctx)

		assert.NoError(t, err)
		assert.Equal(t, filepath.Join(dir, "path"), abs)
	})

	t.Run("it should return an error if the path is not absolute and cwd is not set", func(t *testing.T) {
		f := file.NewFile("path", []byte{})
		ctx := context.New()

		_, err := f.Abs(ctx)

		assert.Error(t, err)
	})
}

func TestFile_Rel(t *testing.T) {
	t.Run("it should return the relative path of the file with context cwd", func(t *testing.T) {
		f := file.NewFile("path", []byte{})
		dir := t.TempDir()
		ctx := context.New()
		ctx = ctx.WithCwd(dir)

		rel, err := f.Rel(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "path", rel)
	})

	t.Run("it should return the relative path of the file with context cwd if the path of the file is absolute", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "path")
		f := file.NewFile(path, []byte{})
		ctx := context.New()
		ctx = ctx.WithCwd(dir)

		rel, err := f.Rel(ctx)

		assert.NoError(t, err)
		assert.Equal(t, "path", rel)
	})

	t.Run("it should return an error if the path of the file is absolute and the cwd is not set", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "path")
		f := file.NewFile(path, []byte{})
		ctx := context.New()

		_, err := f.Rel(ctx)

		assert.Error(t, err)
	})
}

func TestNewFile(t *testing.T) {
	t.Run("it should return a new file", func(t *testing.T) {
		b := []byte("value")
		f := file.NewFile("path", b)

		assert.Implements(t, (*file.File)(nil), f)
		assert.Equal(t, "path", f.Path())
		assert.Equal(t, "value", f.Value())
		assert.Equal(t, b, f.Bytes())
	})
}

func TestNewTextFile(t *testing.T) {
	t.Run("it should return a new file", func(t *testing.T) {
		f := file.NewTextFile("path", "value")

		assert.Implements(t, (*file.File)(nil), f)
		assert.Equal(t, "path", f.Path())
		assert.Equal(t, "value", f.Value())
		assert.Equal(t, []byte("value"), f.Bytes())
	})
}
