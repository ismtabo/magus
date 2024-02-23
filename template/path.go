package template

import (
	"path/filepath"
	"regexp"

	go_errors "github.com/pkg/errors"

	"github.com/ismtabo/magus/v2/context"
)

var (
	pathRegex = regexp.MustCompile(`^[^<>:"|?*\\/\n]+$|^(([a-zA-Z]:)?(\.|\.\.|[^<>:"|?*\\\n]+)?(\\[^<>:"\\|?*]+)+\\?)$|^((\.|\.\.|[^<>:"|?*\\\n]+)?(/[^<>:"|?*\\\n]+)+/?)$`)
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
	// if match := pathRegex.MatchString(value); !match {
	// 	// TODO: Wrap error
	// 	return "", go_errors.Errorf("template returned a non-path path: %s", value)
	// }
	if match := pathRegex.Match([]byte(value)); !match {
		return "", go_errors.Errorf("template returned a non-path path: %s", value)
	}
	if isAbs := filepath.IsAbs(value); isAbs {
		return "", go_errors.Errorf("template returned a non-relative path: %s", value)
	}
	return value, nil
}
