package magic_test

import (
	"testing"

	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/source"
	"github.com/ismtabo/magus/template"
	"github.com/ismtabo/magus/variable"
	"github.com/stretchr/testify/assert"
)

func TestFromManifest(t *testing.T) {
	t.Run("should create a magic from a manifest", func(t *testing.T) {
		mf := manifest.Manifest{
			Version: "1.0.0",
			Name:    "test",
		}

		m := magic.FromManifest(mf)

		assert.Equal(t, magic.NewMagic("1.0.0", "test", []variable.Variable{}, []cast.Cast{}), m)
	})

	t.Run("should create a magic from a manifest with variables", func(t *testing.T) {
		mf := manifest.Manifest{
			Version: "1.0.0",
			Name:    "test",
			Variables: manifest.Variables{
				{
					Name:  "name",
					Value: "value",
				},
			},
		}

		m := magic.FromManifest(mf)

		assert.Equal(t, magic.NewMagic("1.0.0", "test", []variable.Variable{
			variable.NewLiteralVariable("name", "value"),
		}, []cast.Cast{}), m)
	})

	t.Run("should create a magic from a manifest with casts", func(t *testing.T) {
		mf := manifest.Manifest{
			Version: "1.0.0",
			Name:    "test",
			Casts: manifest.Casts{
				"cast": manifest.Cast{
					To:   "to",
					From: "from",
				},
			},
		}

		m := magic.FromManifest(mf)

		assert.Equal(t, magic.NewMagic("1.0.0", "test", []variable.Variable{}, []cast.Cast{
			cast.NewBaseCast(source.NewTemplateSource("from"), template.NewTemplatedString("to"), variable.Variables{}),
		}), m)
	})
}
