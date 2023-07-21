package cast

import (
	"encoding/json"
	"path/filepath"

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
	coll, err := c.collection.Compile(ctx)
	if err != nil {
		return nil, err
	}
	res, ok := c.UnmarshalSlice(coll)
	if !ok {
		return nil, errors.NewValidationError(go_errors.New("collection is not a valid JSON array"))
	}
	files := []file.File{}
	for _, item := range res {
		ctx = ctx.WithVariables(ctx.Variables().Set(c.alias, item))
		dest, err := c.baseCast.dest.Compile(ctx)
		if err != nil {
			// TODO: Wrap error
			return nil, err
		}
		newCwd, err := filepath.Rel(ctx.Cwd(), dest)
		if err != nil {
			// TODO: Wrap error
			return nil, err
		}
		ctx = ctx.WithCwd(newCwd)
		if ok, err := c.filter.Evaluate(ctx); err != nil {
			// TODO: Wrap error
			return nil, err
		} else if ok {
			continue
		} else {
			val, err := c.baseCast.Compile(ctx)
			if err != nil {
				// TODO: Wrap error
				return nil, err
			}
			files = append(files, val...)
		}
	}
	return files, nil
}

func (c *CollectionCast) UnmarshalSlice(coll string) ([]any, bool) {
	var res []any
	err := json.Unmarshal([]byte(coll), &res)
	if err != nil {
		return []any{}, false
	}
	return res, true
}
