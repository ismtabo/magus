package source

import (
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/template"
)

var _ Source = &templateSource{}

type TemplateSource interface {
	Source
}

type templateSource struct {
	template string
}

func NewTemplateSource(template string) TemplateSource {
	return &templateSource{template: template}
}

func (s *templateSource) Compile(ctx context.Context, dest string) ([]file.File, error) {
	value, err := template.Engine.Render(ctx, s.template)
	if err != nil {
		return []file.File{}, err
	}
	return []file.File{
		file.NewTextFile(dest, value),
	}, nil
}
