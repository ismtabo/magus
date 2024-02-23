package fs

import (
	go_errors "errors"
	"os"
	"path"
	"path/filepath"

	"github.com/gookit/goutil/fsutil"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	cp "github.com/otiai10/copy"
)

func ReadFile(ctx context.Context, p string) (file.File, error) {
	if !path.IsAbs(p) {
		p = filepath.Join(ctx.Cwd(), p)
	}
	data, err := os.ReadFile(p)
	if err != nil {
		// TODO: Wrap error
		return nil, err
	}
	return file.NewFile(p, data), nil
}

type ReadDirOptions struct {
	NoFailOnMissing bool
}

func ReadDir(ctx context.Context, p string, opts ReadDirOptions) ([]file.File, error) {
	if !path.IsAbs(p) {
		p = filepath.Join(ctx.Cwd(), p)
	}
	files, err := os.ReadDir(p)
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
			files, err := ReadDir(ctx, filepath.Join(p, f.Name()), opts)
			if err != nil {
				// TODO: Wrap error
				return nil, err
			}
			result = append(result, files...)
			continue
		}
		data, err := os.ReadFile(filepath.Join(p, f.Name()))
		if err != nil {
			// TODO: Wrap error
			return nil, err
		}
		fpath := filepath.Join(p, f.Name())
		result = append(result, file.NewFile(fpath, data))
	}
	return result, nil
}

type ReadFilesOptions struct {
	NoFailOnMissing bool
}

func ReadFiles(ctx context.Context, files []file.File, opts ReadFilesOptions) ([]file.File, error) {
	read_files := []file.File{}
	for _, f := range files {
		path, _ := f.Abs(ctx)
		data, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) && opts.NoFailOnMissing {
				continue
			}
			return nil, err
		}
		read_files = append(read_files, file.NewFile(path, data))
	}
	return read_files, nil
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
