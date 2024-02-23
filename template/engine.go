package template

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/errors"

	"github.com/Masterminds/sprig/v3"
	pluralize "github.com/gertd/go-pluralize"
)

const name = "magus/template"

var (
	pluralizer                   = pluralize.NewClient()
	Engine        TemplateEngine = NewTemplateEngine()
	default_funcs                = template.FuncMap{
		"snake":    strcase.ToSnake,
		"constant": strcase.ToScreamingSnake,
		"pascal":   strcase.ToCamel,
		"camel":    strcase.ToLowerCamel,
		"kebab":    strcase.ToKebab,
		"pluralize": func(s string) string {
			return pluralizer.Plural(s)
		},
		"singularize": func(s string) string {
			return pluralizer.Singular(s)
		},
	}
	funcs = template.FuncMap{}
)

func init() {
	for k, v := range default_funcs {
		funcs[k] = v
	}
	for k, v := range sprig.TxtFuncMap() {
		funcs[k] = v
	}
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
