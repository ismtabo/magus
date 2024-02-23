package cast

import (
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
)

// Cast is the interface that wraps the Compile method.
type Cast interface {
	// Compile compiles the cast.
	Compile(ctx context.Context) ([]file.File, error)
}
