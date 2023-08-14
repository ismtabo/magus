package source

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
)

type Source interface {
	Compile(ctx context.Context, dest string) ([]file.File, error)
}
