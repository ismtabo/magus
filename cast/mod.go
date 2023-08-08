package cast

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
)

// Cast is the interface that wraps the Compile method.
type Cast interface {
	// Compile compiles the cast.
	Compile(ctx context.Context) ([]file.File, error)
}
