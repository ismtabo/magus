package cast

import (
	"encoding/json"

	go_errors "errors"

	"github.com/ismtabo/magus/condition"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/errors"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/template"
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
	for _, item := range res {
		itFiles, err := c.compileItem(ctx, item)
		if err != nil {
			return nil, err
		}
		files = append(files, itFiles...)
	}
	return files, nil
}

func (c *CollectionCast) compileItem(ctx context.Context, item any) ([]file.File, error) {
	ctx = ctx.WithVariable(c.alias, item)
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
