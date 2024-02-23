package cast

import (
	"encoding/json"

	go_errors "errors"

	"github.com/ismtabo/magus/v2/condition"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/errors"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/template"
)

var _ Cast = (*CollectionCast)(nil)

// CollectionCast is a cast that renders files from a source to a destination for a given collection.
type CollectionCast struct {
	collection template.TemplatedString
	alias      string
	filter     condition.Condition
	baseCast   *BaseCast
}

// NewCollectionCast creates a new collection cast.
func NewCollectionCast(coll template.TemplatedString, alias string, filter condition.Condition, baseCast *BaseCast) *CollectionCast {
	return &CollectionCast{
		collection: coll,
		alias:      alias,
		filter:     filter,
		baseCast:   baseCast,
	}
}

// Compile compiles the cast.
func (c *CollectionCast) Compile(ctx context.Context) ([]file.File, error) {
	coll, err := c.collection.Render(ctx)
	if err != nil {
		return nil, err
	}
	res, ok := c.unmarshalSlice(coll)
	if !ok {
		return nil, errors.NewValidationError(go_errors.New("collection is not a valid JSON array"))
	}
	files := []file.File{}
	for idx, value := range res {
		iterCtx := ctx.WithVariable(c.alias, value)
		iterCtx = iterCtx.WithVariable("Index", idx)
		iterCtx = iterCtx.WithVariable("First", idx == 0)
		iterCtx = iterCtx.WithVariable("Last", idx == len(res)-1)
		itFiles, err := c.compileItem(iterCtx)
		if err != nil {
			return nil, err
		}
		files = append(files, itFiles...)
	}
	return files, nil
}

func (c *CollectionCast) compileItem(ctx context.Context) ([]file.File, error) {
	if ok, err := c.filter.Evaluate(ctx); err != nil {
		return nil, err
	} else if !ok {
		return []file.File{}, nil
	}
	return c.baseCast.Compile(ctx)
}

func (c *CollectionCast) unmarshalSlice(coll string) ([]any, bool) {
	var res []any
	err := json.Unmarshal([]byte(coll), &res)
	if err != nil {
		return []any{}, false
	}
	return res, true
}
