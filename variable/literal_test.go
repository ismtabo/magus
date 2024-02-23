package variable_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/variable"
	"github.com/stretchr/testify/assert"
)

func TestLiteralVariable_Name(t *testing.T) {
	t.Run("it should return the name of the variable", func(t *testing.T) {
		v := variable.NewLiteralVariable("name", "value")

		actual := v.Name()

		assert.Equal(t, "name", actual)
	})
}

func TestLiteralVariable_Value(t *testing.T) {
	t.Run("it should return the value of the variable", func(t *testing.T) {
		v := variable.NewLiteralVariable("name", "value")
		actual, err := v.Value(nil)

		assert.NoError(t, err)
		assert.Equal(t, "value", actual)
	})
}
