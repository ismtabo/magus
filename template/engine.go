package template

import (
	"encoding/json"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/errors"
)

const name = "magus/template"

var Engine TemplateEngine = NewTemplateEngine()
var default_funcs = template.FuncMap{
	"lower":    strings.ToLower,
	"upper":    strings.ToUpper,
	"snake":    strcase.ToSnake,
	"constant": strcase.ToScreamingSnake,
	"pascal":   strcase.ToCamel,
	"camel":    strcase.ToLowerCamel,
	"kebab":    strcase.ToKebab,
	"to_json":  ToJson,
}

func ToJson(param any) (string, error) {
	bytes, err := json.Marshal(param)
	if err != nil {
		// TODO: Wrap error
		return "", err
	}
	return string(bytes), nil
}

type TemplateEngine interface {
	// Render renders the given template using the context variables and helpers.
	Render(ctx context.Context, template string) (string, error)
}

type templateEngine struct {
	tmpl *template.Template
}

// NewTemplateEngine returns a new template engine.
func NewTemplateEngine() TemplateEngine {
	return &templateEngine{
		tmpl: template.New(name),
	}
}

// Render renders the given template using the context variables and helpers.
func (e *templateEngine) Render(ctx context.Context, tmplStr string) (string, error) {
	funcs := template.FuncMap{}
	for k, v := range default_funcs {
		funcs[k] = v
	}
	it := ctx.Helpers().Iterator()
	for !it.Done() {
		key, val, _ := it.Next()
		funcs[key] = val
	}
	data := map[string]any{}
	it = ctx.Variables().Iterator()
	for !it.Done() {
		key, val, _ := it.Next()
		data[key] = val
	}
	tmpl, err := e.tmpl.Funcs(funcs).Parse(tmplStr)
	if err != nil {
		return "", errors.NewValidationError(err)
	}
	str := &strings.Builder{}
	err = tmpl.Execute(str, data)
	if err != nil {
		return "", errors.NewRenderError(err)
	}
	return str.String(), nil
}
