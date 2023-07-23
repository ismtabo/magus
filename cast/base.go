package cast

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/source"
	"github.com/ismtabo/magus/template"
	"github.com/ismtabo/magus/variable"
)

var _ Cast = (*BaseCast)(nil)

// BaseCast is a cast that renders files from a source to a destination.
type BaseCast struct {
	src  source.Source
	dest template.TemplatedString
	vars variable.Variables
}

// NewBaseCast creates a new base cast.
func NewBaseCast(src source.Source, dest template.TemplatedString, vars variable.Variables) *BaseCast {
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
	ctx = ctx.WithCwd(dest)
	return c.src.Compile(ctx)
}
