package manifest

import (
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/fs"
)

func Unmarshal(ctx context.Context, path string, manifest *Manifest) error {
	file, err := fs.ReadFile(ctx, path)
	if err != nil {
		return err
	}
	if err := UnmarshalYAML(ctx, file.Bytes(), manifest); err != nil {
		return err
	}
	manifest.File = file
	return nil
}
