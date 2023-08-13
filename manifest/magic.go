package manifest

import (
	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/imports"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/variable"
)

func NewMagic(ctx context.Context, mf Manifest) magic.Magic {
	ctx = imports.WithCtx(ctx)
	if err := imports.Ctx(ctx).Add(ctx, imports.NewImport(CtxManifest(ctx), mf.File)); err != nil {
		panic(err)
	}
	ctx = WithCtxManifest(ctx, &mf)
	cwd, _ := mf.Dir(ctx)
	ctx = ctx.WithCwd(cwd)
	vv := variable.Variables{}
	if mf.Variables != nil {
		for _, v := range mf.Variables {
			vv = append(vv, NewVariable(v))
		}
	}
	casts := []cast.Cast{}
	if mf.Casts != nil {
		for _, c := range mf.Casts {
			casts = append(casts, NewCast(ctx, c))
		}
	}
	return magic.NewMagic(mf.Version, mf.Name, vv, casts)
}
