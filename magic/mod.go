package magic

import (
	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/variable"
)

func FromManifest(mf manifest.Manifest) Magic {
	variables := variable.Variables{}
	if mf.Variables != nil {
		for _, v := range mf.Variables {
			variables = append(variables, variable.NewVariable(v))
		}
	}
	casts := []cast.Cast{}
	if mf.Casts != nil {
		for _, v := range mf.Casts {
			casts = append(casts, cast.NewCast(v))
		}
	}
	return NewMagic(mf.Version, mf.Name, variables, casts)
}
