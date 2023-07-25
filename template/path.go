package template

import (
	"regexp"

	go_errors "github.com/pkg/errors"

	"github.com/ismtabo/magus/context"
)

var (
	pathRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9_-]+|\.{1,2})(?:(?:/|\\)(?:[a-zA-Z0-9_-]+|\.{1,2}))*$`)
)

type TemplatedPath struct {
	string
}

func NewTemplatedPath(path string) TemplatedPath {
	return TemplatedPath{path}
}

func (p TemplatedPath) Render(ctx context.Context) (string, error) {
	value, err := Engine.Render(ctx, p.string)
	if err != nil {
		// TODO: Wrap error
		return "", err
	}
	if match := pathRegex.MatchString(value); !match {
		// TODO: Wrap error
		return "", go_errors.Errorf("template returned a non-path path: %s", value)
	}
	return value, nil
}
