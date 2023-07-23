package manifest_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/manifest"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	t.Run("it should unmarshal a manifest", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(`---
version: 0.1.0
name: magus
root: .
`)
		fp := filepath.Join(t.TempDir(), "manifest.yaml")
		if err := os.WriteFile(fp, y, fs.FileMode(0644)); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()

		err := manifest.Unmarshal(ctx, fp, m)

		assert.NoError(t, err)
		assert.Equal(t, "0.1.0", m.Version)
		assert.Equal(t, "magus", m.Name)
		assert.Equal(t, ".", m.Root)
		assert.Nil(t, m.Variables)
		assert.Nil(t, m.Casts)
	})

	t.Run("it should return an error if the manifest does not exist", func(t *testing.T) {
		ctx := context.New()
		m := &manifest.Manifest{}
		err := manifest.Unmarshal(ctx, "testdata/does-not-exist.yaml", m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the manifest is invalid", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(`---
version: {}
`)
		fp := filepath.Join(t.TempDir(), "manifest.yaml")
		if err := os.WriteFile(fp, y, fs.FileMode(0644)); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()

		err := manifest.Unmarshal(ctx, fp, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the manifest is empty", func(t *testing.T) {
		m := &manifest.Manifest{}
		fp := filepath.Join(t.TempDir(), "manifest.yaml")
		if err := os.WriteFile(fp, []byte(``), fs.FileMode(0644)); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()

		err := manifest.Unmarshal(ctx, fp, m)

		assert.Error(t, err)
	})
}
