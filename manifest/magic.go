package manifest

import (
	"github.com/ismtabo/magus/cast"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/variable"
)

func NewMagic(mf Manifest) magic.Magic {
	variables := variable.Variables{}
	if mf.Variables != nil {
		for _, v := range mf.Variables {
			variables = append(variables, NewVariable(v))
		}
	}
	casts := []cast.Cast{}
	if mf.Casts != nil {
		for _, v := range mf.Casts {
			casts = append(casts, NewCast(v))
		}
	}
	return magic.NewMagic(mf.Version, mf.Name, variables, casts)
}
