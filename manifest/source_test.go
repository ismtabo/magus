package manifest_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/magic"
	"github.com/ismtabo/magus/v2/manifest"
	"github.com/ismtabo/magus/v2/source"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestNewSource(t *testing.T) {
	t.Run("it should return a template source if the source is a string", func(t *testing.T) {
		src := manifest.NewSource(context.New(), manifest.Source{}.FromString("Hello World!\n"))

		assert.Implements(t, (*source.Source)(nil), src)
	})

	t.Run("it should return a magic source if the source is a map", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "magic.yml")
		data := []byte(dedent.Dedent(`
		version: "1"
		name: "magic"
		root: "."

		casts:
		  hello:
		    to: "world"
		    from: "Hello World!"
		`))
		if err := os.WriteFile(path, data, 0644); err != nil {
			t.Fatal(err)
		}
		src := manifest.NewSource(context.New().WithCwd("."), manifest.Source{}.FromStruct(
			manifest.MagicSource{
				Magic: path,
			},
		))

		assert.Implements(t, (*source.Source)(nil), src)
		assert.Implements(t, (*magic.Magic)(nil), src)
	})

	t.Run("it should return a magic source relative to the context cwd", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "magic.yml")
		data := []byte(dedent.Dedent(`
		version: "1"
		name: "magic"
		root: "."
		casts:
		  hello:
		    to: "world"
		    from: "Hello World!"
		`))
		if err := os.WriteFile(path, data, 0644); err != nil {
			t.Fatal(err)
		}
		ctx := context.New().WithCwd(dir)
		src := manifest.NewSource(ctx, manifest.Source{}.FromStruct(
			manifest.MagicSource{
				Magic: "magic.yml",
			},
		))

		assert.Implements(t, (*source.Source)(nil), src)
		assert.Implements(t, (*magic.Magic)(nil), src)
	})

	t.Run("it should panic if the source is a magic but it does not point to a file", func(t *testing.T) {
		assert.Panics(t, func() {
			manifest.NewSource(context.New(), manifest.Source{}.FromStruct(
				manifest.MagicSource{
					Magic: "hello",
				}))
		})
	})
}
