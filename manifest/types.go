package manifest

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

// EitherStringOrStruct is a union type of string or struct.
type EitherStringOrStruct[T any] struct {
	Str    string `validate:"required_without=Struct"`
	Struct T      `validate:"required_without=Str"`
	is_map bool
}

func (e *EitherStringOrStruct[T]) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind == yaml.ScalarNode {
		return node.Decode(&e.Str)
	}
	if err := node.Decode(&e.Struct); err != nil {
		return err
	}
	e.is_map = true
	return nil
}

func (e *EitherStringOrStruct[T]) IsString() bool {
	return reflect.ValueOf(e.Struct).IsZero()
}

func (e EitherStringOrStruct[T]) FromString(s string) EitherStringOrStruct[T] {
	return EitherStringOrStruct[T]{
		Str:    s,
		is_map: false,
	}
}

func (e EitherStringOrStruct[T]) FromStruct(s T) EitherStringOrStruct[T] {
	return EitherStringOrStruct[T]{
		Struct: s,
		is_map: true,
	}
}
