package context

import (
	go_context "context"

	"github.com/benbjohnson/immutable"
)

type Context interface {
	go_context.Context
	Cwd() string
	Variables() *immutable.Map[string, any]
	Helpers() *immutable.Map[string, any]
	WithVariables(vals *immutable.Map[string, any]) Context
	WithHelpers(helpers *immutable.Map[string, any]) Context
	WithCwd(cwd string) Context
}

type context struct {
	go_context.Context
	cwd       string
	variables *immutable.Map[string, any]
	helpers   *immutable.Map[string, any]
}

func New() Context {
	return &context{
		Context:   go_context.Background(),
		cwd:       ".",
		variables: immutable.NewMap[string, any](nil),
		helpers:   immutable.NewMap[string, any](nil),
	}
}

func (ctx *context) Cwd() string {
	return ctx.cwd
}

func (ctx *context) Variables() *immutable.Map[string, any] {
	return ctx.variables
}

func (ctx *context) Helpers() *immutable.Map[string, any] {
	return ctx.helpers
}

func (ctx *context) WithVariables(variables *immutable.Map[string, any]) Context {
	newVars := ctx.variables
	it := variables.Iterator()
	for !it.Done() {
		k, v, _ := it.Next()
		newVars = newVars.Set(k, v)
	}
	return &context{
		Context:   ctx,
		cwd:       ctx.Cwd(),
		variables: newVars,
		helpers:   ctx.Helpers(),
	}
}

func (ctx *context) WithHelpers(helpers *immutable.Map[string, any]) Context {
	newHelpers := ctx.helpers
	it := helpers.Iterator()
	for !it.Done() {
		k, v, _ := it.Next()
		newHelpers = newHelpers.Set(k, v)
	}
	return &context{
		Context:   ctx,
		cwd:       ctx.Cwd(),
		variables: ctx.Variables(),
		helpers:   helpers,
	}
}

func (ctx *context) WithCwd(cwd string) Context {
	return &context{
		Context:   ctx,
		cwd:       cwd,
		variables: ctx.Variables(),
		helpers:   ctx.Helpers(),
	}
}
