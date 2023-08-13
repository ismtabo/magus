package manifest

import "github.com/ismtabo/magus/context"

type contextKey = string

const (
	currentManifestKey contextKey = "currentManifest"
)

func WithCtxManifest(ctx context.Context, manifest *Manifest) context.Context {
	return context.WithValue(ctx, currentManifestKey, manifest)
}

func CtxManifest(ctx context.Context) *Manifest {
	if value := ctx.Value(currentManifestKey); value != nil {
		return value.(*Manifest)
	}
	return nil
}
