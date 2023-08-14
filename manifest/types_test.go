package manifest_test

import (
	"testing"

	"github.com/ismtabo/magus/manifest"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type TestStruct struct {
	Str string `yaml:"str"`
}

func TestEitherStringOrMap_UnmarshalYAML(t *testing.T) {
	t.Run("it should unmarshal a string", func(t *testing.T) {
		var m manifest.EitherStringOrStruct[TestStruct]
		if err := yaml.Unmarshal([]byte("test"), &m); err != nil {
			t.Fatal(err)
		}
		assert.True(t, m.IsString())
		assert.Equal(t, "test", m.Str)
	})

	t.Run("it should unmarshal a map", func(t *testing.T) {
		var m manifest.EitherStringOrStruct[TestStruct]
		if err := yaml.Unmarshal([]byte("str: test"), &m); err != nil {
			t.Fatal(err)
		}
		assert.False(t, m.IsString())
		assert.Equal(t, TestStruct{
			Str: "test",
		}, m.Struct)
	})

	t.Run("it should fail if the value is not a string nor a map", func(t *testing.T) {
		var m manifest.EitherStringOrStruct[TestStruct]
		if err := yaml.Unmarshal([]byte("[1, 2, 3]"), &m); err == nil {
			t.Fatal("expected an error")
		}
	})
}

func TestEitherStringOrMap_IsString(t *testing.T) {
	t.Run("it should return true if the value is a string", func(t *testing.T) {
		m := manifest.EitherStringOrStruct[TestStruct]{
			Str: "test",
		}
		assert.True(t, m.IsString())
	})

	t.Run("it should return false if the value is a map", func(t *testing.T) {
		m := manifest.EitherStringOrStruct[TestStruct]{
			Struct: TestStruct{
				Str: "test",
			},
		}
		assert.False(t, m.IsString())
	})
}

func TestEitherStringOrMap_FromString(t *testing.T) {
	t.Run("it should create a string", func(t *testing.T) {
		m := manifest.EitherStringOrStruct[TestStruct]{}.FromString("test")
		assert.True(t, m.IsString())
		assert.Equal(t, "test", m.Str)
	})
}

func TestEitherStringOrMap_FromMap(t *testing.T) {
	t.Run("it should create a map", func(t *testing.T) {
		m := manifest.EitherStringOrStruct[TestStruct]{}.FromStruct(
			TestStruct{
				Str: "test",
			},
		)
		assert.False(t, m.IsString())
		assert.Equal(t, TestStruct{
			Str: "test",
		}, m.Struct)
	})
}
