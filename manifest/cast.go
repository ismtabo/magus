package manifest

import (
	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/template"
)

// NewCast creates a new cast from a manifest cast.
func NewCast(c Cast) cast.Cast {
	dest := template.NewTemplatedPath(c.To)
	src := NewSource(c.From)
	vars := NewVariables(c.Variables)
	hasCond := c.If != "" || c.Unless != ""
	hasEach := c.Each != ""
	hasEitherCondOrEach := hasCond || hasEach
	baseCast := cast.NewBaseCast(src, dest, vars)
	if !hasEitherCondOrEach {
		return baseCast
	}
	var cond condition.Condition = condition.NewAlwaysTrueCondition()
	if c.If != "" {
		cond = condition.NewTemplateCondition(c.If)
	}
	if c.Unless != "" {
		cond = condition.NewNegatedCondition(condition.NewTemplateCondition(c.Unless))
	}
	condCast := cast.NewConditionalCast(cond, baseCast)
	if !hasEach {
		return condCast
	}
	each := template.NewTemplatedString(c.Each)
	as := "It"
	if c.As != "" {
		as = c.As
	}
	var filter condition.Condition = condition.NewAlwaysTrueCondition()
	if c.Include != "" {
		filter = condition.NewTemplateCondition(c.Include)
	}
	if c.Omit != "" {
		filter = condition.NewNegatedCondition(condition.NewTemplateCondition(c.Omit))
	}
	collCast := cast.NewCollectionCast(each, as, filter, baseCast)
	if !hasCond {
		return collCast
	}
	return cast.NewConditionalCollectionCast(condCast, collCast)
}
