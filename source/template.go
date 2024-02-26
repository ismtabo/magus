package source

import (
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/file"
	"github.com/ismtabo/magus/v2/template"
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
	if s.template == "" {
		return []file.File{
			file.NewTextFile(dest, ""),
		}, nil
	}
	value, err := template.Engine.Render(ctx, s.template)
	if err != nil {
		return []file.File{}, err
	}
	return []file.File{
		file.NewTextFile(dest, value),
	}, nil
}
