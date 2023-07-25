package fs

import (
	go_errors "errors"
	"os"
	"path/filepath"

	"github.com/gookit/goutil/fsutil"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	cp "github.com/otiai10/copy"
)

func ReadFile(ctx context.Context, path string) (file.File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// TODO: Wrap error
		return nil, err
	}
	return file.NewFile(path, data), nil
}

type ReadDirOptions struct {
	NoFailOnMissing bool
}

func ReadDir(ctx context.Context, path string, opts ReadDirOptions) ([]file.File, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) && opts.NoFailOnMissing {
			return []file.File{}, nil
		}
		return nil, err
	}
	// NOTE: options are not propagated to recursive calls
	opts = ReadDirOptions{}
	var result []file.File
	for _, f := range files {
		if f.IsDir() {
			files, err := ReadDir(ctx, filepath.Join(path, f.Name()), opts)
			if err != nil {
				// TODO: Wrap error
				return nil, err
			}
			result = append(result, files...)
			continue
		}
		data, err := os.ReadFile(filepath.Join(path, f.Name()))
		if err != nil {
			// TODO: Wrap error
			return nil, err
		}
		fpath := filepath.Join(path, f.Name())
		result = append(result, file.NewFile(fpath, data))
	}
	return result, nil
}

func WriteFiles(ctx context.Context, files []file.File) error {
	for _, f := range files {
		path, _ := f.Abs(ctx)
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := os.WriteFile(path, f.Bytes(), 0644); err != nil {
			// TODO: Wrap error
			return err
		}
	}
	return nil
}

func CopyDir(ctx context.Context, src, dest string) error {
	if !filepath.IsAbs(src) {
		if ctx.Cwd() == "" {
			// TODO: Wrap error
			return go_errors.New("context cwd is empty")
		}
		src = filepath.Join(ctx.Cwd(), src)
	}
	if !filepath.IsAbs(dest) {
		if ctx.Cwd() == "" {
			// TODO: Wrap error
			return go_errors.New("context cwd is empty")
		}
		dest = filepath.Join(ctx.Cwd(), dest)
	}
	if err := cp.Copy(src, dest); err != nil {
		// TODO: Wrap error
		return err
	}
	return nil
}

func CleanDir(ctx context.Context, path string) error {
	if !filepath.IsAbs(path) {
		if ctx.Cwd() == "" {
			// TODO: Wrap error
			return go_errors.New("context cwd is empty")
		}
		path = filepath.Join(ctx.Cwd(), path)
	}
	if err := fsutil.RemoveSub(path); err != nil {
		// TODO: Wrap error
		return err
	}
	return nil
}
