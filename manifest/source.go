package manifest

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/source"
)

func NewSource(ctx context.Context, s Source) source.Source {
	if s.IsString() {
		return source.NewTemplateSource(s.Str)
	}
	return newSourceFromMagicMap(ctx, s.Struct)
}

func newSourceFromMagicMap(ctx context.Context, ms MagicSource) source.Source {
	m := Manifest{}
	if err := Unmarshal(ctx, ms.Magic, &m); err != nil {
		panic(err)
	}
	return NewMagic(ctx, m)
}
