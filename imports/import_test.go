package imports_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/imports"
	"github.com/stretchr/testify/assert"
)

func TestNewImport(t *testing.T) {
	t.Run(`it should return an Import with the given "from" and "to" files`, func(t *testing.T) {
		from := file.NewFile("from", nil)
		to := file.NewFile("to", nil)
		i := imports.NewImport(from, to)
		assert.Equal(t, from, i.From())
		assert.Equal(t, to, i.To())
	})
}
