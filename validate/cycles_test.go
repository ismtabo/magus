package validate_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/manifest"
	"github.com/ismtabo/magus/v2/validate"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestValidateNoCycles(t *testing.T) {
	t.Run("it should return nil if there are magic casts", func(t *testing.T) {
		mf := manifest.Manifest{
			File: file.NewFile("manifest.yaml", nil),
			Casts: map[string]manifest.Cast{
				"foo": {
					From: manifest.Source{}.FromString("foo"),
				},
			},
		}
		assert.NoError(t, validate.NoCycles(context.New(), mf))
	})

	t.Run("it should return nil if there are no cycles", func(t *testing.T) {
		cwd := t.TempDir()
		ctx := context.New().WithCwd(cwd)
		mf1_content := dedent.Dedent(`
		---
		version: 1
		name: foo
		root: .
		casts:
		  bar:
		    to: bar
		    from:
		      magic: bar.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "foo.yaml"), []byte(mf1_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf2_content := dedent.Dedent(`
		---
		version: 1
		name: bar
		root: .
		casts:
		  baz:
		    to: baz
		    from: string
		`)
		if err := os.WriteFile(filepath.Join(cwd, "bar.yaml"), []byte(mf2_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf := manifest.Manifest{
			File: file.NewFile("manifest.yaml", nil),
			Casts: map[string]manifest.Cast{
				"foo": {
					To: "foo",
					From: manifest.Source{}.FromStruct(manifest.MagicSource{
						Magic: "foo.yaml",
					}),
				},
			},
		}
		assert.NoError(t, validate.NoCycles(ctx, mf))
	})

	t.Run("it should return err if magic cast file does not exists", func(t *testing.T) {
		cwd := t.TempDir()
		ctx := context.New().WithCwd(cwd)
		mf := manifest.Manifest{
			File: file.NewFile("manifest.yaml", nil),
			Casts: map[string]manifest.Cast{
				"foo": {
					To: "foo",
					From: manifest.Source{}.FromStruct(manifest.MagicSource{
						Magic: "unknown.yaml",
					}),
				},
			},
		}

		err := validate.NoCycles(ctx, mf)
		assert.Error(t, err)
	})

	t.Run("it should return err if magic uses itself", func(t *testing.T) {
		cwd := t.TempDir()
		ctx := context.New().WithCwd(cwd)
		mf_content := dedent.Dedent(`
		---
		version: 1
		name: manifest
		root: .
		casts:
		  foo:
		    to: bar
		    from:
		      magic: manifest.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "manifest.yaml"), []byte(mf_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf := manifest.Manifest{
			File: file.NewFile("manifest.yaml", nil),
			Casts: map[string]manifest.Cast{
				"foo": {
					To: "foo",
					From: manifest.Source{}.FromStruct(manifest.MagicSource{
						Magic: "manifest.yaml",
					}),
				},
			},
		}

		err := validate.NoCycles(ctx, mf)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("already imported %s", filepath.Join(cwd, "manifest.yaml")))
	})

	t.Run("it should return err if child magic cast uses parent", func(t *testing.T) {
		cwd := t.TempDir()
		ctx := context.New().WithCwd(cwd)
		mf_content := dedent.Dedent(`
		---
		version: 1
		name: manifest
		root: .
		casts:
		  foo:
		    to: foo
		    from:
		      magic: foo.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "manifest.yaml"), []byte(mf_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf1_content := dedent.Dedent(`
		---
		version: 1
		name: foo
		root: .
		casts:
		  bar:
		    to: bar
		    from:
		      magic: manifest.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "foo.yaml"), []byte(mf1_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf := manifest.Manifest{
			File: file.NewFile("manifest.yaml", nil),
			Casts: map[string]manifest.Cast{
				"foo": {
					To: "foo",
					From: manifest.Source{}.FromStruct(manifest.MagicSource{
						Magic: "foo.yaml",
					}),
				},
			},
		}

		err := validate.NoCycles(ctx, mf)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("already imported %s", filepath.Join(cwd, "manifest.yaml")))
	})

	t.Run("it should return err if grandchild magic cast uses parent", func(t *testing.T) {
		cwd := t.TempDir()
		ctx := context.New().WithCwd(cwd)
		mf_content := dedent.Dedent(`
		---
		version: 1
		name: manifest
		root: .
		casts:
		  foo:
		    to: foo
		    from:
		      magic: foo.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "manifest.yaml"), []byte(mf_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf1_content := dedent.Dedent(`
		---
		version: 1
		name: foo
		root: .
		casts:
		  bar:
		    to: bar
		    from:
		      magic: bar.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "foo.yaml"), []byte(mf1_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf2_content := dedent.Dedent(`
		---
		version: 1
		name: bar
		root: .
		casts:
		  baz:
		    to: baz
		    from:
		      magic: manifest.yaml
		`)
		if err := os.WriteFile(filepath.Join(cwd, "bar.yaml"), []byte(mf2_content), 0644); err != nil {
			t.Fatal(err)
		}
		mf := manifest.Manifest{
			File: file.NewFile("manifest.yaml", nil),
			Casts: map[string]manifest.Cast{
				"foo": {
					To: "foo",
					From: manifest.Source{}.FromStruct(manifest.MagicSource{
						Magic: "foo.yaml",
					}),
				},
			},
		}

		err := validate.NoCycles(ctx, mf)
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("already imported %s", filepath.Join(cwd, "manifest.yaml")))
	})
}
