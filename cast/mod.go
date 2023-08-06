package cast

import (
	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/source"
	"github.com/ismtabo/magus/template"
	"github.com/ismtabo/magus/variable"
)

// Cast is the interface that wraps the Compile method.
type Cast interface {
	// Compile compiles the cast.
	Compile(ctx context.Context) ([]file.File, error)
}

// NewCast creates a new cast from a manifest cast.
func NewCast(cast manifest.Cast) Cast {
	dest := template.NewTemplatedPath(cast.To)
	src := source.NewSource(cast.From)
	vars := variable.NewVariables(cast.Variables)
	hasCond := cast.If != "" || cast.Unless != ""
	hasEach := cast.Each != ""
	hasEitherCondOrEach := hasCond || hasEach
	baseCast := NewBaseCast(src, dest, vars)
	if !hasEitherCondOrEach {
		return baseCast
	}
	var cond condition.Condition = condition.NewAlwaysTrueCondition()
	if cast.If != "" {
		cond = condition.NewTemplateCondition(cast.If)
	}
	if cast.Unless != "" {
		cond = condition.NewNegatedCondition(condition.NewTemplateCondition(cast.Unless))
	}
	condCast := NewConditionalCast(cond, baseCast)
	if !hasEach {
		return condCast
	}
	each := template.NewTemplatedString(cast.Each)
	as := "It"
	if cast.As != "" {
		as = cast.As
	}
	var filter condition.Condition = condition.NewAlwaysTrueCondition()
	if cast.Include != "" {
		filter = condition.NewTemplateCondition(cast.Include)
	}
	if cast.Omit != "" {
		filter = condition.NewNegatedCondition(condition.NewTemplateCondition(cast.Omit))
	}
	collCast := NewCollectionCast(each, as, filter, baseCast)
	if !hasCond {
		return collCast
	}
	return NewConditionalCollectionCast(condCast, collCast)
}
