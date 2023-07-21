package source

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/manifest"
)

type Source interface {
	Compile(ctx context.Context) ([]file.File, error)
}

func NewSource(source manifest.Source) Source {
	val, ok := source.(string)
	if ok {
		return NewTemplateSource(val)
	}
	panic("not implemented")
}
