package cast

import (
	"path/filepath"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/source"
	"github.com/ismtabo/magus/v2/template"
	"github.com/ismtabo/magus/v2/variable"
)

var _ Cast = (*BaseCast)(nil)

// BaseCast is a cast that renders files from a source to a destination.
type BaseCast struct {
	src  source.Source
	dest template.TemplatedPath
	vars variable.Variables
}

// NewBaseCast creates a new base cast.
func NewBaseCast(src source.Source, dest template.TemplatedPath, vars variable.Variables) *BaseCast {
	return &BaseCast{
		src:  src,
		dest: dest,
		vars: vars,
	}
}

// Compile compiles the cast.
func (c *BaseCast) Compile(ctx context.Context) ([]file.File, error) {
	for _, v := range c.vars {
		value, err := v.Value(ctx)
		if err != nil {
			// TODO: Wrap error
			return nil, err
		}
		ctx = ctx.WithVariable(v.Name(), value)
	}
	dest, err := c.dest.Render(ctx)
	if err != nil {
		// TODO: Wrap error
		return nil, err
	}
	dest = filepath.Join(ctx.Cwd(), dest)
	ctx = ctx.WithCwd(filepath.Dir(dest))
	return c.src.Compile(ctx, dest)
}
