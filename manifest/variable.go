package manifest

import "github.com/ismtabo/magus/v2/variable"

// New creates a new variable.
func NewVariable(v Variable) variable.Variable {
	name := v.Name
	if v.Env != "" {
		return variable.NewEnvironmentVariable(name, v.Env)
	}
	if v.Template != "" {
		return variable.NewTemplateVariable(name, v.Template)
	}
	return variable.NewLiteralVariable(name, v.Value)
}

// NewVariables creates a new slice of variables.
func NewVariables(vv []Variable) variable.Variables {
	vars := make(variable.Variables, len(vv))
	for i, variable := range vv {
		vars[i] = NewVariable(variable)
	}
	return vars
}
