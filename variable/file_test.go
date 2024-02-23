package variable_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/variable"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestFromFile(t *testing.T) {
	t.Run("it should return the variables from a yaml file", func(t *testing.T) {
		cwd := t.TempDir()
		path := filepath.Join(cwd, "vars.yaml")
		data := []byte(dedent.Dedent(`
		foo: bar
		baz: qux
		`))
		if err := os.WriteFile(path, data, 0755); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		expected := variable.Variables{
			variable.NewLiteralVariable("foo", "bar"),
			variable.NewLiteralVariable("baz", "qux"),
		}

		actual, err := variable.FromFile(ctx, path)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("it should return the variables from a json file", func(t *testing.T) {
		cwd := t.TempDir()
		path := filepath.Join(cwd, "vars.json")
		data := []byte(dedent.Dedent(`{ "foo": "bar", "baz": "qux" }`))
		if err := os.WriteFile(path, data, 0755); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)
		expected := variable.Variables{
			variable.NewLiteralVariable("foo", "bar"),
			variable.NewLiteralVariable("baz", "qux"),
		}

		actual, err := variable.FromFile(ctx, path)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("it should return an error if the file is not a yaml file or a json file", func(t *testing.T) {
		cwd := t.TempDir()
		path := filepath.Join(cwd, "vars.txt")
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)

		_, err := variable.FromFile(ctx, path)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the file does not exist", func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithCwd("testdata")
		path := "vars-does-not-exist.yaml"

		_, err := variable.FromFile(ctx, path)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the file is not a valid yaml file", func(t *testing.T) {
		cwd := t.TempDir()
		path := filepath.Join(cwd, "invalid-vars.yaml")
		data := []byte(dedent.Dedent(`
			foo: "
		`))
		if err := os.WriteFile(path, data, 0755); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)

		_, err := variable.FromFile(ctx, path)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the file is not a valid json file", func(t *testing.T) {
		cwd := t.TempDir()
		path := filepath.Join(cwd, "invalid-vars.json")
		data := []byte(dedent.Dedent(`
		{
			"foo": ""
		`))
		if err := os.WriteFile(path, data, 0755); err != nil {
			t.Fatal(err)
		}
		ctx := context.New()
		ctx = ctx.WithCwd(cwd)

		_, err := variable.FromFile(ctx, path)

		assert.Error(t, err)
	})
}
