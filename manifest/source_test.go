package manifest_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/source"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestNewSource(t *testing.T) {
	t.Run("it should return a template source if the source is a string", func(t *testing.T) {
		src := manifest.NewSource("Hello World!\n")

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
		src := manifest.NewSource(map[string]interface{}{
			"magic": path,
		})

		assert.Implements(t, (*source.Source)(nil), src)
	})

	t.Run("it should panic if the source is not a string neither a map", func(t *testing.T) {
		assert.Panics(t, func() {
			manifest.NewSource(123)
		})
	})

	t.Run("it should panic if the source is a map but it does not contain a magic key", func(t *testing.T) {
		assert.Panics(t, func() {
			manifest.NewSource(map[string]interface{}{
				"hello": "world",
			})
		})
	})

	t.Run("it should panic if the source is a map but the magic key is not a string", func(t *testing.T) {
		assert.Panics(t, func() {
			manifest.NewSource(map[string]interface{}{
				"magic": 123,
			})
		})
	})

	t.Run("it should panic if the source is a map but the magic key is an empty string", func(t *testing.T) {
		assert.Panics(t, func() {
			manifest.NewSource(map[string]interface{}{
				"magic": "",
			})
		})
	})

	t.Run("it should panic if the source is a map but the magic key is a string that does not point to a file", func(t *testing.T) {
		assert.Panics(t, func() {
			manifest.NewSource(map[string]interface{}{
				"magic": "hello",
			})
		})
	})
}
