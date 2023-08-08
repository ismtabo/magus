package manifest_test

import (
	"testing"

	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/manifest"
	"github.com/stretchr/testify/assert"
)

func TestNewCast(t *testing.T) {
	t.Run("it should return a new base cast", func(t *testing.T) {
		m := manifest.Cast{
			To:   "/tmp",
			From: "Hello World",
			Variables: manifest.Variables{
				manifest.Variable{
					Name:  "foo",
					Value: "bar",
				},
			},
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.BaseCast{}, c)
	})

	t.Run("it should return a new conditional cast", func(t *testing.T) {
		m := manifest.Cast{
			To:   "/tmp",
			From: "Hello World",
			If:   "true",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.ConditionalCast{}, c)
	})

	t.Run("it should return a new conditional cast with unless", func(t *testing.T) {
		m := manifest.Cast{
			To:     "/tmp",
			From:   "Hello World",
			Unless: "false",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.ConditionalCast{}, c)
	})

	t.Run("it should return a new loop cast", func(t *testing.T) {
		m := manifest.Cast{
			To:   "/tmp",
			From: "Hello World",
			Each: "[]",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.CollectionCast{}, c)
	})

	t.Run("it should return a new loop cast with as", func(t *testing.T) {
		m := manifest.Cast{
			To:   "/tmp",
			From: "Hello World",
			Each: "[]",
			As:   "foo",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.CollectionCast{}, c)
	})

	t.Run("it should return a new loop cast with include", func(t *testing.T) {
		m := manifest.Cast{
			To:      "/tmp",
			From:    "Hello World",
			Each:    "[]",
			Include: "true",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.CollectionCast{}, c)
	})

	t.Run("it should return a new loop cast with omit", func(t *testing.T) {
		m := manifest.Cast{
			To:   "/tmp",
			From: "Hello World",
			Each: "[]",
			Omit: "false",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.CollectionCast{}, c)
	})

	t.Run("it should return a new conditional loop cast", func(t *testing.T) {
		m := manifest.Cast{
			To:   "/tmp",
			From: "Hello World",
			Each: "[]",
			If:   "true",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.ConditionalCollectionCast{}, c)
	})

	t.Run("it should return a new conditional loop cast with unless", func(t *testing.T) {
		m := manifest.Cast{
			To:     "/tmp",
			From:   "Hello World",
			Each:   "[]",
			Unless: "false",
		}

		c := manifest.NewCast(m)

		assert.IsType(t, &cast.ConditionalCollectionCast{}, c)
	})
}
