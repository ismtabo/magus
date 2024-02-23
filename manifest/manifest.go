package manifest

import (
	"github.com/ismtabo/magus/v2/file"
)

// Manifest is the representation of a manifest file.
type Manifest struct {
	// File is the file of the manifest.
	file.File `yaml:"-"`
	// Version is the version of the manifest.
	Version Version `yaml:"version" validate:"required,semver"`
	// Name is the name of the manifest.
	Name string `yaml:"name" validate:"required"`
	// Root is the output root dir of the manifest.
	Root string `yaml:"root" validate:"required,dirpath"`
	// Variables is the variables of the manifest.
	Variables Variables `yaml:"variables" validate:"dive"`
	// Casts is the casts of the manifest.
	Casts Casts `yaml:"casts" validate:"dive"`
}

// Variables is the variables of the manifest.
type Variables = []Variable

// Variable is the variable of the manifest.
type Variable struct {
	// Name is the name of the variable.
	Name string `yaml:"name" validate:"required"`
	// Value is the value of the variable.
	Value interface{} `yaml:"value" validate:"required_without_all=Template Env"`
	// Template is the template of the variable.
	Template string `yaml:"template" validate:"required_without_all=Value Env"`
	// Name is the name of the environment variable.
	Env string `yaml:"env" validate:"required_without_all=Value Template"`
}

// Casts is the casts of the manifest.
type Casts = map[string]Cast

// Cast is the cast of the manifest.
type Cast struct {
	// To is the output path of the cast
	To string `yaml:"to" validate:"required"`
	// From is the input source of the cast
	From Source `yaml:"from" validate:"dive"`
	// Variables is the variables of the cast
	Variables Variables `yaml:"variables" validate:"dive"`
	// If is the condition of the cast
	If string `yaml:"if,omitempty" validate:"excluded_with=Unless"`
	// Unless is the negated condition of the cast
	Unless string `yaml:"unless,omitempty" validate:"excluded_with=If"`
	// Each is the loop of the cast
	Each string `yaml:"each,omitempty" validate:"required_with_any=As Include Omit"`
	// As is the name of the loop variable
	As string `yaml:"as,omitempty"`
	// Include is the condition of the loop
	Include string `yaml:"include,omitempty" validate:"excluded_with=Omit"`
	// Omit is the negated condition of the loop
	Omit string `yaml:"omit,omitempty" validate:"excluded_with=Include"`
}

// Source is the source of the cast.
type Source = EitherStringOrStruct[MagicSource]

// MagicSource is the magic source of the cast.
type MagicSource struct {
	// Magic is the path of the included manifest.
	Magic string `mapstructure:"magic" yaml:"magic" validate:"omitempty,filepath"`
}
