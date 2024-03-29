package validate

import (
	"path/filepath"

	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/imports"
	"github.com/ismtabo/magus/v2/manifest"
)

func NoCycles(ctx context.Context, mf manifest.Manifest) error {
	ctx = imports.WithCtx(ctx)
	svc := imports.Ctx(ctx)
	_ = svc.Add(ctx, imports.NewImport(nil, mf.File))
	dir, _ := mf.Dir(ctx)
	ctx = ctx.WithCwd(dir)
	for _, c := range mf.Casts {
		if c.From.IsString() {
			continue
		}
		cmf := manifest.Manifest{}
		path := filepath.Join(dir, c.From.Struct.Magic)
		if err := svc.Add(ctx, imports.NewImport(mf.File, file.NewFile(path, nil))); err != nil {
			return err
		}
		if err := manifest.Unmarshal(ctx, path, &cmf); err != nil {
			return err
		}
		if err := NoCycles(ctx, cmf); err != nil {
			return err
		}
	}
	return nil
}
