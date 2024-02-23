package imports

import (
	"github.com/benbjohnson/immutable"
	"github.com/ismtabo/magus/v2/context"
)

type importCtxKey = string

const (
	// ImportCtxKey is the key used to store the import context in the
	// context.Context.
	ImportCtxKey importCtxKey = "import"
)

type ImportService interface {
	Add(ctx context.Context, i Import) error
}

type ImportServiceImpl struct {
	imports immutable.Set[string]
}

func WithCtx(ctx context.Context) context.Context {
	if svc := Ctx(ctx); svc != nil {
		return ctx
	}
	return context.WithValue(ctx, ImportCtxKey, NewImportService())
}

func Ctx(ctx context.Context) ImportService {
	if value := ctx.Value(ImportCtxKey); value != nil {
		return value.(ImportService)
	}
	return nil
}

func NewImportService() ImportService {
	return &ImportServiceImpl{
		imports: immutable.NewSet[string](nil),
	}
}

func (s *ImportServiceImpl) Add(ctx context.Context, i Import) error {
	if s == nil {
		return nil
	}
	if i == nil {
		return nil
	}
	if i.To() == nil {
		return nil
	}
	abs_path, _ := i.To().Abs(ctx)
	if s.imports.Has(abs_path) {
		return AlreadyImportedError(i)
	}
	s.imports = s.imports.Add(abs_path)
	return nil
}
