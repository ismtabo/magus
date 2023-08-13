package domain_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/domain"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/variable"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("it should render the magic files", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:    file.NewFile("manifest.yaml", nil),
			Version: "1",
			Name:    "test",
			Root:    ".",
			Variables: []manifest.Variable{
				{
					Name:  "foo",
					Value: "foo",
				},
			},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
						{{ .bar }}
						{{ .baz }}
					`)),
					Variables: []manifest.Variable{
						{
							Name:  "bar",
							Value: "bar",
						},
					},
				},
			},
		}
		opts := domain.GenerateOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("baz", "baz"),
			},
		}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(files))
		assert.Equal(t, filepath.Join(outDir, "magic"), files[0].Path())
		assert.Equal(t, dedent.Dedent(`
			foo
			bar
			baz
		`), files[0].Value())
		assert.FileExists(t, filepath.Join(outDir, "magic"))
	})

	t.Run("it should use manifest root if empty directory is given", func(t *testing.T) {
		cwd := t.TempDir()
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:    file.NewFile("manifest.yaml", nil),
			Version: "1",
			Name:    "test",
			Root:    "out",
			Variables: []manifest.Variable{
				{
					Name:  "foo",
					Value: "foo",
				},
			},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
						{{ .bar }}
						{{ .baz }}
					`)),
					Variables: []manifest.Variable{
						{
							Name:  "bar",
							Value: "bar",
						},
					},
				},
			},
		}
		opts := domain.GenerateOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("baz", "baz"),
			},
		}

		files, err := domain.Generate(ctx, "", mfst, opts)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(files))
		assert.Equal(t, filepath.Join("out", "magic"), files[0].Path())
		assert.Equal(t, dedent.Dedent(`
			foo
			bar
			baz
		`), files[0].Value())
		assert.FileExists(t, filepath.Join(cwd, "out", "magic"))
	})

	t.Run("it should not create files on dry run", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:    file.NewFile("manifest.yaml", nil),
			Version: "1",
			Name:    "test",
			Root:    ".",
			Variables: []manifest.Variable{
				{
					Name:  "foo",
					Value: "foo",
				},
			},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
						{{ .bar }}
						{{ .baz }}
					`)),
					Variables: []manifest.Variable{
						{
							Name:  "bar",
							Value: "bar",
						},
					},
				},
			},
		}
		opts := domain.GenerateOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("baz", "baz"),
			},
			DryRun: true,
		}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.NoError(t, err)
		assert.Len(t, files, 1)
		assert.Equal(t, filepath.Join(outDir, "magic"), files[0].Path())
		assert.Equal(t, dedent.Dedent(`
			foo
			bar
			baz
		`), files[0].Value())
		assert.NoFileExists(t, filepath.Join(outDir, "magic"))
	})

	t.Run("it should not overwrite existing files", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		eContent := "existent"
		if err := os.Mkdir(outDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(outDir, "magic"), []byte(eContent), 0644); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:    file.NewFile("manifest.yaml", nil),
			Version: "1",
			Name:    "test",
			Root:    ".",
			Variables: []manifest.Variable{
				{
					Name:  "foo",
					Value: "foo",
				},
			},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
						{{ .bar }}
						{{ .baz }}
					`)),
					Variables: []manifest.Variable{
						{
							Name:  "bar",
							Value: "bar",
						},
					},
				},
			},
		}
		opts := domain.GenerateOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("baz", "baz"),
			},
		}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.NoError(t, err)
		assert.Empty(t, files)
		assert.FileExists(t, filepath.Join(outDir, "magic"))
		if content, err := os.ReadFile(filepath.Join(outDir, "magic")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, eContent, string(content))
		}
	})

	t.Run("it should overwrite existing files if overwrite is set", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		eContent := "existent"
		if err := os.Mkdir(outDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(outDir, "magic"), []byte(eContent), 0644); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:    file.NewFile("manifest.yaml", nil),
			Version: "1",
			Name:    "test",
			Root:    ".",
			Variables: []manifest.Variable{
				{
					Name:  "foo",
					Value: "foo",
				},
			},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
						{{ .bar }}
						{{ .baz }}
					`)),
					Variables: []manifest.Variable{
						{
							Name:  "bar",
							Value: "bar",
						},
					},
				},
			},
		}
		opts := domain.GenerateOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("baz", "baz"),
			},
			Overwrite: true,
		}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.NoError(t, err)
		assert.Len(t, files, 1)
		assert.Equal(t, filepath.Join(outDir, "magic"), files[0].Path())
		assert.Equal(t, dedent.Dedent(`
			foo
			bar
			baz
		`), files[0].Value())
		assert.FileExists(t, filepath.Join(outDir, "magic"))
		if content, err := os.ReadFile(filepath.Join(outDir, "magic")); err != nil {
			t.Fatal(err)
		} else {
			assert.NotEqual(t, eContent, string(content))
			assert.Equal(t, dedent.Dedent(`
				foo
				bar
				baz
			`), string(content))
		}
	})

	t.Run("it should clean existing files if clean is set", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		eContent := "existent"
		if err := os.Mkdir(outDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(outDir, "magic"), []byte(eContent), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(outDir, "other"), []byte(eContent), 0644); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:    file.NewFile("manifest.yaml", nil),
			Version: "1",
			Name:    "test",
			Root:    ".",
			Variables: []manifest.Variable{
				{
					Name:  "foo",
					Value: "foo",
				},
			},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
						{{ .bar }}
						{{ .baz }}
					`)),
					Variables: []manifest.Variable{
						{
							Name:  "bar",
							Value: "bar",
						},
					},
				},
			},
		}
		opts := domain.GenerateOptions{
			Variables: variable.Variables{
				variable.NewLiteralVariable("baz", "baz"),
			},
			Clean: true,
		}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.NoError(t, err)
		assert.Len(t, files, 1)
		assert.Equal(t, filepath.Join(outDir, "magic"), files[0].Path())
		assert.Equal(t, dedent.Dedent(`
			foo
			bar
			baz
		`), files[0].Value())
		assert.NoFileExists(t, filepath.Join(outDir, "other"))
		assert.FileExists(t, filepath.Join(outDir, "magic"))
		if content, err := os.ReadFile(filepath.Join(outDir, "magic")); err != nil {
			t.Fatal(err)
		} else {
			assert.NotEqual(t, eContent, string(content))
			assert.Equal(t, dedent.Dedent(`
				foo
				bar
				baz
			`), string(content))
		}
	})

	t.Run("it should return error if the magic render fails", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:      file.NewFile("manifest.yaml", nil),
			Version:   "1",
			Name:      "test",
			Root:      ".",
			Variables: []manifest.Variable{},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo
					`)),
					Variables: []manifest.Variable{},
				},
			},
		}
		opts := domain.GenerateOptions{}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.Error(t, err)
		assert.Nil(t, files)
	})

	t.Run("it should return error if there are conflicts between files", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:      file.NewFile("manifest.yaml", nil),
			Version:   "1",
			Name:      "test",
			Root:      ".",
			Variables: []manifest.Variable{},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
					`)),
					Variables: []manifest.Variable{},
				},
				"other": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromString(dedent.Dedent(`
						{{ .foo }}
					`)),
					Variables: []manifest.Variable{},
				},
			},
		}
		opts := domain.GenerateOptions{}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.Error(t, err)
		assert.Nil(t, files)
	})

	t.Run("it should render files from magic casts", func(t *testing.T) {
		cwd := t.TempDir()
		outDir := filepath.Join(cwd, "out")
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		mfst := manifest.Manifest{
			File:      file.NewFile("manifest.yaml", nil),
			Version:   "1",
			Name:      "test",
			Root:      ".",
			Variables: []manifest.Variable{},
			Casts: manifest.Casts{
				"magic": manifest.Cast{
					To: "magic",
					From: manifest.Source{}.FromStruct(
						manifest.MagicSource{
							Magic: "foo.yaml",
						},
					),
					Variables: []manifest.Variable{},
				},
			},
		}
		opts := domain.GenerateOptions{}
		foo_mfst := dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		    from: baz
		`)
		if err := os.WriteFile(filepath.Join(cwd, "foo.yaml"), []byte(foo_mfst), 0644); err != nil {
			t.Fatal(err)
		}

		files, err := domain.Generate(ctx, outDir, mfst, opts)

		assert.NoError(t, err)
		assert.Len(t, files, 1)
		assert.Equal(t, filepath.Join(outDir, "magic", "bar"), files[0].Path())
		assert.Equal(t, "baz", files[0].Value())
		assert.FileExists(t, filepath.Join(outDir, "magic", "bar"))
	})
}
