package manifest_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/cast"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/imports"
	"github.com/ismtabo/magus/v2/magic"
	"github.com/ismtabo/magus/v2/manifest"
	"github.com/ismtabo/magus/v2/source"
	"github.com/ismtabo/magus/v2/template"
	"github.com/ismtabo/magus/v2/variable"
	"github.com/stretchr/testify/assert"
)

func TestFromManifest(t *testing.T) {
	t.Run("it should create a magic from a manifest", func(t *testing.T) {
		ctx := context.New().WithCwd(t.TempDir())
		mf := manifest.Manifest{
			File: file.NewFile("test", nil),
			Name: "test",
		}

		m := manifest.NewMagic(ctx, mf)

		assert.Equal(t, magic.NewMagic("test", []variable.Variable{}, []cast.Cast{}), m)
	})

	t.Run("it should create a magic from a manifest with variables", func(t *testing.T) {
		ctx := context.New().WithCwd(t.TempDir())
		mf := manifest.Manifest{
			File: file.NewFile("test", nil),
			Name: "test",
			Variables: manifest.Variables{
				{
					Name:  "name",
					Value: "value",
				},
			},
		}

		m := manifest.NewMagic(ctx, mf)

		assert.Equal(t, magic.NewMagic("test", []variable.Variable{
			variable.NewLiteralVariable("name", "value"),
		}, []cast.Cast{}), m)
	})

	t.Run("it should create a magic from a manifest with casts", func(t *testing.T) {
		ctx := context.New().WithCwd(t.TempDir())
		ctx = imports.WithCtx(ctx)
		f := file.NewFile("test", nil)
		mf := manifest.Manifest{
			File: f,
			Name: "test",
			Casts: manifest.Casts{
				"cast": manifest.Cast{
					To:   "to",
					From: manifest.Source{}.FromString("from"),
				},
			},
		}

		m := manifest.NewMagic(ctx, mf)

		assert.Equal(t, magic.NewMagic("test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("from"), template.NewTemplatedPath("to"), variable.Variables{}),
		}), m)
	})

	t.Run("it should panic if the manifest is already imported in the context", func(t *testing.T) {
		ctx := context.New().WithCwd(t.TempDir())
		ctx = imports.WithCtx(ctx)
		svc := imports.Ctx(ctx)
		f := file.NewFile("test", nil)
		if err := svc.Add(ctx, imports.NewImport(nil, f)); err != nil {
			t.Fatal(err)
		}
		mf := manifest.Manifest{
			File: f,
			Name: "test",
		}
		assert.Panics(t, func() {
			manifest.NewMagic(ctx, mf)
		})
	})
}
