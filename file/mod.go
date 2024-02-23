package file

import (
	"path/filepath"

	go_errors "errors"

	"github.com/ismtabo/magus/v2/context"
)

type File interface {
	// Path is the path of the file.
	Path() string
	// Dir is the directory of the file.
	Dir(ctx context.Context) (string, error)
	// Abs returns the absolute path of the file.
	Abs(ctx context.Context) (string, error)
	// Rel returns the relative path of the file with context cwd.
	Rel(ctx context.Context) (string, error)
	// Value is the content of the file.
	Value() string
	// Bytes is the content of the file.
	Bytes() []byte
}

type file struct {
	path string
	data []byte
}

// NewFile creates a new file.
func NewFile(path string, data []byte) File {
	return &file{
		path: path,
		data: data,
	}
}

func NewTextFile(path string, data string) File {
	return NewFile(path, []byte(data))
}

// Path is the path of the file.
func (f file) Path() string {
	return f.path
}

// Dir is the directory of the file.
func (f file) Dir(ctx context.Context) (string, error) {
	path, err := f.Abs(ctx)
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}

// Value is the content of the file.
func (f file) Value() string {
	return string(f.data)
}

// Bytes is the content of the file.
func (f file) Bytes() []byte {
	return f.data
}

// Abs returns the absolute path of the file.
func (f file) Abs(ctx context.Context) (string, error) {
	if filepath.IsAbs(f.path) {
		return f.path, nil
	}
	if ctx.Cwd() == "" {
		return "", go_errors.New("context cwd is empty")
	}
	return filepath.Join(ctx.Cwd(), f.path), nil
}

// Rel returns the relative path of the file with context cwd.
func (f file) Rel(ctx context.Context) (string, error) {
	if !filepath.IsAbs(f.path) {
		return f.path, nil
	}
	if ctx.Cwd() == "" {
		return "", go_errors.New("context cwd is empty")
	}
	path, err := filepath.Rel(ctx.Cwd(), f.path)
	if err != nil {
		// TODO: Wrap error
		return "", err
	}
	return path, nil
}
