package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/fs"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	t.Run("should return a file", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "file.txt")
		if err := os.WriteFile(p, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}

		f, err := fs.ReadFile(ctx, p)

		assert.NoError(t, err)
		assert.Equal(t, p, f.Path())
		assert.Equal(t, []byte("Hello, World!"), f.Bytes())
	})

	t.Run("should return an error if the file does not exist", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "file.txt")

		_, err := fs.ReadFile(ctx, p)

		assert.Error(t, err)
	})

	t.Run("should return an error if the path is a directory", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "dir")
		if err := os.Mkdir(p, 0755); err != nil {
			t.Fatal(err)
		}

		_, err := fs.ReadFile(ctx, p)

		assert.Error(t, err)
	})
}

func TestReadDir(t *testing.T) {
	t.Run("should return a list of files", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		subdir := filepath.Join(dir, "subdir")
		p2 := filepath.Join(subdir, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.Mkdir(subdir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		opts := fs.ReadDirOptions{}

		files, err := fs.ReadDir(ctx, dir, opts)

		assert.NoError(t, err)
		assert.Len(t, files, 2)
		assert.Equal(t, p1, files[0].Path())
		assert.Equal(t, p2, files[1].Path())
	})

	t.Run("should return an empty list if the directory does not exist and option NoFailOnMissing is true", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "dir")
		opts := fs.ReadDirOptions{
			NoFailOnMissing: true,
		}

		files, err := fs.ReadDir(ctx, p, opts)

		assert.NoError(t, err)
		assert.Empty(t, files)
	})

	t.Run("should return an error if the directory does not exist", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "dir")
		opts := fs.ReadDirOptions{}

		_, err := fs.ReadDir(ctx, p, opts)

		assert.Error(t, err)
	})

	t.Run("should return an error if the path is a file", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "file.txt")
		if err := os.WriteFile(p, []byte("Hello World!"), 0755); err != nil {
			t.Fatal(err)
		}
		opts := fs.ReadDirOptions{}

		_, err := fs.ReadDir(ctx, p, opts)

		assert.Error(t, err)
	})
}

func TestReadFiles(t *testing.T) {
	t.Run("should return a list of files", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		files := []file.File{
			file.NewFile(p1, []byte("Hello, World!")),
			file.NewFile(p2, []byte("Hello, World!")),
		}

		read_files, err := fs.ReadFiles(ctx, files, fs.ReadFilesOptions{})

		assert.NoError(t, err)
		assert.Len(t, read_files, 2)
		assert.Equal(t, p1, read_files[0].Path())
		assert.Equal(t, p2, read_files[1].Path())
	})

	t.Run("should return a list of files even if some files do not exist and option NoFailOnMissing is true", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		files := []file.File{
			file.NewFile(p1, []byte("Hello, World!")),
			file.NewFile(p2, []byte("Hello, World!")),
		}
		opts := fs.ReadFilesOptions{
			NoFailOnMissing: true,
		}

		read_files, err := fs.ReadFiles(ctx, files, opts)

		assert.NoError(t, err)
		assert.Len(t, read_files, 1)
		assert.Equal(t, p1, read_files[0].Path())
	})

	t.Run("should return an error if some files do not exist", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		files := []file.File{
			file.NewFile(p1, []byte("Hello, World!")),
			file.NewFile(p2, []byte("Hello, World!")),
		}

		_, err := fs.ReadFiles(ctx, files, fs.ReadFilesOptions{})

		assert.Error(t, err)
	})
}

func TestWriteFiles(t *testing.T) {
	t.Run("should write a list of files", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		files := []file.File{
			file.NewFile(p1, []byte("Hello, World!")),
			file.NewFile(p2, []byte("Hello, World!")),
		}

		err := fs.WriteFiles(ctx, files)

		assert.NoError(t, err)
		assert.FileExists(t, p1)
		assert.FileExists(t, p2)
	})

	t.Run("should write a list of files with relative path to cwd", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		ctx = ctx.WithCwd(dir)
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		files := []file.File{
			file.NewFile("file1.txt", []byte("Hello, World!")),
			file.NewFile("file2.txt", []byte("Hello, World!")),
		}

		err := fs.WriteFiles(ctx, files)

		assert.NoError(t, err)
		assert.FileExists(t, p1)
		assert.FileExists(t, p2)
	})

	t.Run("should write a list of files even if files parent directories does not exists", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		files := []file.File{
			file.NewFile(p1, []byte("Hello, World!")),
			file.NewFile(p2, []byte("Hello, World!")),
		}

		err := fs.WriteFiles(ctx, files)

		assert.NoError(t, err)
		assert.FileExists(t, p1)
		assert.FileExists(t, p2)
	})

	t.Run("should return an error if the file cannot be written", func(t *testing.T) {
		ctx := context.New()
		dir := filepath.Join(t.TempDir(), "subdir")
		p := filepath.Join(dir, "file.txt")
		if err := os.Mkdir(dir, 0555); err != nil {
			t.Fatal(err)
		}
		files := []file.File{
			file.NewFile(p, []byte("Hello, World!")),
		}

		err := fs.WriteFiles(ctx, files)

		assert.Error(t, err)
	})

	t.Run("should return an error if the file path is not absolute and cwd is not set", func(t *testing.T) {
		ctx := context.New()
		files := []file.File{
			file.NewFile("file.txt", []byte("Hello, World!")),
		}

		err := fs.WriteFiles(ctx, files)

		assert.Error(t, err)
	})

	t.Run("should return an error if the file path is a directory", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		files := []file.File{
			file.NewFile(dir, []byte("Hello, World!")),
		}

		err := fs.WriteFiles(ctx, files)

		assert.Error(t, err)
	})
}

func TestCopyDir(t *testing.T) {
	t.Run("should copy a directory", func(t *testing.T) {
		ctx := context.New()
		src := t.TempDir()
		p1 := filepath.Join(src, "file1.txt")
		p2 := filepath.Join(src, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		dst := t.TempDir()

		err := fs.CopyDir(ctx, src, dst)

		assert.NoError(t, err)
		assert.FileExists(t, filepath.Join(dst, "file1.txt"))
		assert.FileExists(t, filepath.Join(dst, "file2.txt"))
	})

	t.Run("should copy a directory if the src path is relative", func(t *testing.T) {
		ctx := context.New()
		src := t.TempDir()
		p1 := filepath.Join(src, "file1.txt")
		p2 := filepath.Join(src, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		dst := t.TempDir()
		ctx = ctx.WithCwd(src)

		err := fs.CopyDir(ctx, ".", dst)

		assert.NoError(t, err)
		assert.FileExists(t, filepath.Join(dst, "file1.txt"))
		assert.FileExists(t, filepath.Join(dst, "file2.txt"))
	})

	t.Run("should copy a directory if the dest path is relative", func(t *testing.T) {
		ctx := context.New()
		src := t.TempDir()
		p1 := filepath.Join(src, "file1.txt")
		p2 := filepath.Join(src, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		dst := t.TempDir()
		ctx = ctx.WithCwd(dst)

		err := fs.CopyDir(ctx, src, ".")

		assert.NoError(t, err)
		assert.FileExists(t, filepath.Join(dst, "file1.txt"))
		assert.FileExists(t, filepath.Join(dst, "file2.txt"))
	})

	t.Run("should return an error if the source directory does not exist", func(t *testing.T) {
		ctx := context.New()
		src := filepath.Join(t.TempDir(), "src")
		dst := t.TempDir()

		err := fs.CopyDir(ctx, src, dst)

		assert.Error(t, err)
	})

	t.Run("should return an error if the source path is a file", func(t *testing.T) {
		ctx := context.New()
		src := filepath.Join(t.TempDir(), "file.txt")
		if err := os.WriteFile(src, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		dst := t.TempDir()

		err := fs.CopyDir(ctx, src, dst)

		assert.Error(t, err)
	})
}

func TestCleanDir(t *testing.T) {
	t.Run("should remove all files from a directory", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}

		err := fs.CleanDir(ctx, dir)

		assert.NoError(t, err)
		assert.NoFileExists(t, p1)
		assert.NoFileExists(t, p2)
	})

	t.Run("should remove all files from a directory if the path is relative", func(t *testing.T) {
		ctx := context.New()
		dir := t.TempDir()
		ctx = ctx.WithCwd(dir)
		p1 := filepath.Join(dir, "file1.txt")
		p2 := filepath.Join(dir, "file2.txt")
		if err := os.WriteFile(p1, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p2, []byte("Hello, World!"), 0644); err != nil {
			t.Fatal(err)
		}

		err := fs.CleanDir(ctx, ".")

		assert.NoError(t, err)
		assert.NoFileExists(t, p1)
		assert.NoFileExists(t, p2)
	})

	t.Run("should not return an error if the directory does not exist", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "dir")

		err := fs.CleanDir(ctx, p)

		assert.NoError(t, err)
	})

	t.Run("should not return an error if the path is a file", func(t *testing.T) {
		ctx := context.New()
		p := filepath.Join(t.TempDir(), "file.txt")
		if err := os.WriteFile(p, []byte("Hello World!"), 0755); err != nil {
			t.Fatal(err)
		}

		err := fs.CleanDir(ctx, p)

		assert.NoError(t, err)
	})

	t.Run("should return an error if the path is relative and cwd is not set", func(t *testing.T) {
		ctx := context.New()

		err := fs.CleanDir(ctx, ".")

		assert.Error(t, err)
	})

	t.Run("should return an error if the path is a directory and cannot be removed", func(t *testing.T) {
		ctx := context.New()
		dir := filepath.Join(t.TempDir(), "dir")
		f := filepath.Join(dir, "file.txt")
		if err := os.Mkdir(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(f, []byte("Hello World!"), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.Chmod(dir, 0555); err != nil {
			t.Fatal(err)
		}
		defer os.Chmod(dir, 0755)

		err := fs.CleanDir(ctx, dir)

		assert.Error(t, err)
	})
}
