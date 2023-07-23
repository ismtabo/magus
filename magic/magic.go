package magic

import (
	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/variable"
)

type Magic interface {
	Render(ctx context.Context) ([]file.File, error)
}

type magic struct {
	version string
	name    string
	vars    []variable.Variable
	casts   []cast.Cast
}

func NewMagic(version, name string, vars []variable.Variable, casts []cast.Cast) Magic {
	return &magic{
		version: version,
		name:    name,
		vars:    vars,
		casts:   casts,
	}
}

func (m *magic) Render(ctx context.Context) ([]file.File, error) {
	files := []file.File{}
	for _, v := range m.vars {
		value, err := v.Value(ctx)
		if err != nil {
			// TODO: Wrap error
			return nil, err
		}
		ctx = ctx.WithVariable(v.Name(), value)
	}
	for _, c := range m.casts {
		result, err := c.Compile(ctx)
		if err != nil {
			return []file.File{}, err
		}
		files = append(files, result...)
	}
	return files, nil
}
