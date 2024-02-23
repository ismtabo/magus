package manifest_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/manifest"
	"github.com/stretchr/testify/assert"
)

func TestWithCtxManifest(t *testing.T) {
	t.Run(`it should return a new context with the given Manifest`, func(t *testing.T) {
		mf := manifest.Manifest{}
		ctx := manifest.WithCtxManifest(context.New(), &mf)
		assert.NotNil(t, ctx)
		assert.Implements(t, (*context.Context)(nil), ctx)
		assert.Equal(t, &mf, manifest.CtxManifest(ctx))
	})
	t.Run(`it should return a new context with the updated Manifest`, func(t *testing.T) {
		mf1 := manifest.Manifest{}
		ctx := manifest.WithCtxManifest(context.New(), &mf1)
		mf2 := manifest.Manifest{}
		ctx = manifest.WithCtxManifest(ctx, &mf2)
		assert.NotNil(t, ctx)
		assert.Implements(t, (*context.Context)(nil), ctx)
		assert.Equal(t, &mf2, manifest.CtxManifest(ctx))
	})
}

func TestCtxManifest(t *testing.T) {
	t.Run(`it should return the Manifest stored in the context`, func(t *testing.T) {
		mf := manifest.Manifest{}
		ctx := manifest.WithCtxManifest(context.New(), &mf)
		assert.NotNil(t, manifest.CtxManifest(ctx))
		assert.Equal(t, &mf, manifest.CtxManifest(ctx))
	})
	t.Run(`it should return nil if the context does not contain a Manifest`, func(t *testing.T) {
		assert.Nil(t, manifest.CtxManifest(context.New()))
	})
}
