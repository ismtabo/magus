package imports_test

import (
	"testing"

	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/imports"
	"github.com/stretchr/testify/assert"
)

func TestAlreadyImportedError(t *testing.T) {
	t.Run(`it should return an error with the message "already imported <path>"`, func(t *testing.T) {
		i := imports.NewImport(nil, file.NewFile("file", nil))
		err := imports.AlreadyImportedError(i)
		assert.Error(t, err)
		assert.EqualError(t, err, "already imported file")
	})
}
