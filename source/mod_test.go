package source_test

import (
	"testing"

	"github.com/ismtabo/magus/source"
	"github.com/stretchr/testify/assert"
)

func TestNewSource(t *testing.T) {
	t.Run("it should return a template source if the source is a string", func(t *testing.T) {
		src := source.NewSource("Hello World!\n")

		assert.Implements(t, (*source.Source)(nil), src)
	})

	t.Run("it should panic if the source is not a string", func(t *testing.T) {
		assert.Panics(t, func() {
			source.NewSource(123)
		})
	})
}
