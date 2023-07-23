package variable_test

import (
	"testing"

	"github.com/ismtabo/magus/variable"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentVariable_Name(t *testing.T) {
	t.Run("it should return the name of the variable", func(t *testing.T) {
		v := variable.NewEnvironmentVariable("name", "env")

		actual := v.Name()

		assert.Equal(t, "name", actual)
	})
}

func TestEnvironmentVariable_Value(t *testing.T) {
	t.Run("it should return the value of the environment variable", func(t *testing.T) {
		t.Setenv("env", "value")
		v := variable.NewEnvironmentVariable("name", "env")
		actual, err := v.Value(nil)

		assert.NoError(t, err)
		assert.Equal(t, "value", actual)
	})
}
