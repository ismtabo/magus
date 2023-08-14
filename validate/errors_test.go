package validate_test

import (
	"testing"

	"github.com/ismtabo/magus/validate"
	"github.com/stretchr/testify/assert"
)

func TestWrapInvalidVersion(t *testing.T) {
	err := validate.WrapInvalidVersionErrorf("foo: %s", "bar")
	assert.Equal(t, "invalid version: foo: bar", err.Error())
}
