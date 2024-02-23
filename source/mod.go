package source

import (
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
)

type Source interface {
	Compile(ctx context.Context, dest string) ([]file.File, error)
}
