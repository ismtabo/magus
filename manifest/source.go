package manifest

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/source"
	"github.com/mitchellh/mapstructure"
)

type magicSource struct {
	Magic string `mapstructure:"magic" validate:"required"`
}

func NewSource(s Source) source.Source {
	val, ok := s.(string)
	if ok {
		return source.NewTemplateSource(val)
	}
	val_map, ok := s.(map[any]any)
	if !ok {
		panic("not implemented")
	}
	ms := magicSource{}
	if err := mapstructure.Decode(val_map, &ms); err != nil {
		panic(err)
	}
	if err := v.Struct(ms); err != nil {
		panic(err)
	}
	return newSourceFromMagicMap(context.New(), ms)
}

func newSourceFromMagicMap(ctx context.Context, ms magicSource) source.Source {
	m := Manifest{}
	if err := Unmarshal(ctx, ms.Magic, &m); err != nil {
		panic(err)
	}
	return NewMagic(m)
}
