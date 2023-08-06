package manifest_test

import (
	"testing"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/manifest"
	"github.com/lithammer/dedent"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalYAML(t *testing.T) {
	t.Run("it should unmarshal a manifest", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.NoError(t, err)
		assert.Equal(t, "0.1.0", m.Version)
		assert.Equal(t, "magus", m.Name)
		assert.Equal(t, ".", m.Root)
		assert.Nil(t, m.Variables)
		assert.Nil(t, m.Casts)
	})

	t.Run("it should return an error if the manifest is invalid", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(`---
		version: {}
		`)
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if the manifest is empty", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(``)
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should marshal variables", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		variables:
		  - name: foo
		    value: bar
		  - name: baz
		    template: "{{ .foo }}"
		  - name: qux
		    env: QUX
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.NoError(t, err)
		assert.Equal(t, "0.1.0", m.Version)
		assert.Equal(t, "magus", m.Name)
		assert.Equal(t, ".", m.Root)
		assert.Equal(t, 3, len(m.Variables))
		assert.Equal(t, "foo", m.Variables[0].Name)
		assert.Equal(t, "bar", m.Variables[0].Value)
		assert.Equal(t, "baz", m.Variables[1].Name)
		assert.Equal(t, "{{ .foo }}", m.Variables[1].Template)
		assert.Equal(t, "qux", m.Variables[2].Name)
		assert.Equal(t, "QUX", m.Variables[2].Env)
	})

	t.Run("it should marshal casts", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  base:
		    to: foo
		    from: bar
		  with-variables:
		    to: foo
		    from: bar
		    variables:
		      - name: foo
		        value: bar
		      - name: baz
		        template: "{{ .foo }}"
		      - name: qux
		        env: QUX
		  with-if:
		    to: foo
		    from: bar
		    if: "{{ .foo }}"
		  with-unless:
		    to: foo
		    from: bar
		    unless: "{{ .bar }}"
		  with-each:
		    to: foo
		    from: bar
		    each: "{{ .qux }}"
		  with-as:
		    to: foo
		    from: bar
		    each: "{{ .qux }}"
		    as: quux
		  with-include:
		    to: foo
		    from: bar
		    each: "{{ .qux }}"
		    include: "{{ .quuz }}"
		  with-omit:
		    to: foo
		    from: bar
		    each: "{{ .qux }}"
		    omit: "{{ .corge }}"
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.NoError(t, err)
		assert.Equal(t, "0.1.0", m.Version)
		assert.Equal(t, "magus", m.Name)
		assert.Equal(t, ".", m.Root)
		assert.Equal(t, 8, len(m.Casts))
		assert.Equal(t, "foo", m.Casts["base"].To)
		assert.Equal(t, "bar", m.Casts["base"].From)
		assert.Equal(t, "foo", m.Casts["with-variables"].To)
		assert.Equal(t, "bar", m.Casts["with-variables"].From)
		assert.Equal(t, 3, len(m.Casts["with-variables"].Variables))
		assert.Equal(t, "foo", m.Casts["with-variables"].Variables[0].Name)
		assert.Equal(t, "bar", m.Casts["with-variables"].Variables[0].Value)
		assert.Equal(t, "baz", m.Casts["with-variables"].Variables[1].Name)
		assert.Equal(t, "{{ .foo }}", m.Casts["with-variables"].Variables[1].Template)
		assert.Equal(t, "qux", m.Casts["with-variables"].Variables[2].Name)
		assert.Equal(t, "QUX", m.Casts["with-variables"].Variables[2].Env)
		assert.Equal(t, "{{ .foo }}", m.Casts["with-if"].If)
		assert.Equal(t, "{{ .bar }}", m.Casts["with-unless"].Unless)
		assert.Equal(t, "{{ .qux }}", m.Casts["with-each"].Each)
		assert.Equal(t, "quux", m.Casts["with-as"].As)
		assert.Equal(t, "{{ .quuz }}", m.Casts["with-include"].Include)
		assert.Equal(t, "{{ .corge }}", m.Casts["with-omit"].Omit)
	})

	t.Run("it should return an error if a cast with missing to", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    from: bar
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if a cast with missing from", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if a cast with both present if and unless", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		    from: baz
		    if: "{{ .foo }}"
		    unless: "{{ .bar }}"
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if missing each in a cast with as", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		    from: baz
		    as: quux
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if missing each in a cast with include", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		    from: baz
		    include: quux
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if missing each in a cast with omit", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		    from: baz
		    omit: quux
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})

	t.Run("it should return an error if a cast with both present include and omit", func(t *testing.T) {
		m := &manifest.Manifest{}
		y := []byte(dedent.Dedent(`
		---
		version: 0.1.0
		name: magus
		root: .
		casts:
		  foo:
		    to: bar
		    from: baz
		    each: "{{ .qux }}"
		    include: "{{ .quuz }}"
		    omit: "{{ .corge }}"
		`))
		ctx := context.New()

		err := manifest.UnmarshalYAML(ctx, y, m)

		assert.Error(t, err)
	})
}
