package manifest

// Manifest is the representation of a manifest file.
type Manifest struct {
	// Version is the version of the manifest.
	Version string `json:"version" yaml:"version"`
	// Name is the name of the manifest.
	Name string `json:"name" yaml:"name"`
	// Root is the output root dir of the manifest.
	Root string `json:"root" yaml:"root"`
	// Variables is the variables of the manifest.
	Variables Variables `json:"variables" yaml:"variables"`
}

// Variables is the variables of the manifest.
type Variables = []Variable

// Variable is the variable of the manifest.
type Variable struct {
	// Name is the name of the variable.
	Name string `json:"name" yaml:"name"`
	// Value is the value of the variable.
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
	// Template is the template of the variable.
	Template string `json:"template,omitempty" yaml:"template,omitempty"`
	// Name is the name of the environment variable.
	Env string `json:"env,omitempty" yaml:"env,omitempty"`
}

// Casts is the casts of the manifest.
type Casts = map[string]Cast

// Cast is the cast of the manifest.
type Cast struct {
	// To is the output path of the cast
	To string `json:"to" yaml:"to"`
	// From is the input source of the cast
	From Source `json:"from" yaml:"from"`
	// Variables is the variables of the cast
	Variables Variables `json:"variables" yaml:"variables"`
	// If is the condition of the cast
	If string `json:"if,omitempty" yaml:"if,omitempty"`
	// Unless is the negated condition of the cast
	Unless string `json:"unless,omitempty" yaml:"unless,omitempty"`
	// Each is the loop of the cast
	Each string `json:"each,omitempty" yaml:"each,omitempty"`
	// As is the name of the loop variable
	As string `json:"as,omitempty" yaml:"as,omitempty"`
	// Include is the condition of the loop
	Include string `json:"include,omitempty" yaml:"include,omitempty"`
	// Omit is the negated condition of the loop
	Omit string `json:"exclude,omitempty" yaml:"exclude,omitempty"`
}

// Source is the source of the cast.
type Source = interface{}
