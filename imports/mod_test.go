package imports_test

import (
	"testing"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/imports"
	"github.com/stretchr/testify/assert"
)

func TestWithCtx(t *testing.T) {
	t.Run(`it should return a new context with the given Import`, func(t *testing.T) {
		ctx := imports.WithCtx(context.New())
		assert.NotNil(t, ctx)
		assert.Implements(t, (*context.Context)(nil), ctx)
	})
	t.Run(`it should return the same context if the context already contains an ImportService`, func(t *testing.T) {
		ctx := imports.WithCtx(context.New())
		ctx = imports.WithCtx(ctx)
		assert.NotNil(t, ctx)
		assert.Implements(t, (*context.Context)(nil), ctx)
	})
}

func TestCtx(t *testing.T) {
	t.Run(`it should return the ImportService stored in the context`, func(t *testing.T) {
		ctx := imports.WithCtx(context.New())
		assert.NotNil(t, imports.Ctx(ctx))
		assert.Implements(t, (*imports.ImportService)(nil), imports.Ctx(ctx))
	})
	t.Run(`it should return nil if the context does not contain an ImportService`, func(t *testing.T) {
		assert.Nil(t, imports.Ctx(context.New()))
	})
}

func TestNewImportService(t *testing.T) {
	t.Run(`it should return a new ImportService`, func(t *testing.T) {
		assert.NotNil(t, imports.NewImportService())
		assert.Implements(t, (*imports.ImportService)(nil), imports.NewImportService())
	})
}

func TestImportServiceImpl_Add(t *testing.T) {
	t.Run(`it should return nil if the ImportService is nil`, func(t *testing.T) {
		assert.Nil(t, (*imports.ImportServiceImpl)(nil).Add(nil, nil))
	})
	t.Run(`it should return nil if the Import is nil`, func(t *testing.T) {
		s := imports.NewImportService()
		assert.Nil(t, s.Add(nil, nil))
	})
	t.Run(`it should return nil if the Import is not already imported`, func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithCwd(t.TempDir())
		s := imports.NewImportService()
		from := file.NewFile("from", nil)
		to := file.NewFile("to", nil)
		i := imports.NewImport(from, to)

		assert.Nil(t, s.Add(ctx, i))
	})
	t.Run(`it should return an error if the Import is already imported`, func(t *testing.T) {
		ctx := context.New()
		ctx = ctx.WithCwd(t.TempDir())
		s := imports.NewImportService()
		to := file.NewFile("file", nil)
		i := imports.NewImport(nil, to)
		assert.NoError(t, s.Add(ctx, i))
		assert.EqualError(t, s.Add(ctx, i), imports.AlreadyImportedError(i).Error())
	})
	t.Run(`it should return nil if the Import's To is nil`, func(t *testing.T) {
		s := imports.NewImportService()
		i := imports.NewImport(nil, nil)
		assert.Nil(t, s.Add(nil, i))
	})
}
