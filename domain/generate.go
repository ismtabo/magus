package domain

import (
	"os"
	"path/filepath"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/filetree"
	"github.com/ismtabo/magus/fs"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/manifest"
	"github.com/samber/lo"
)

type GenerateOptions struct {
	DryRun    bool
	Overwrite bool
	Clean     bool
}

func Generate(ctx context.Context, dest string, mfst manifest.Manifest, opts GenerateOptions) ([]file.File, error) {
	ctx = ctx.WithCwd(filepath.Join(ctx.Cwd(), mfst.Root))
	mgc := magic.FromManifest(mfst)

	files, err := mgc.Render(ctx.WithCwd(dest))
	if err != nil {
		return nil, err
	}

	if err := filetree.AssertNotHaveWriteConflicts(ctx, files); err != nil {
		return nil, err
	}

	if !opts.Overwrite && !opts.Clean {
		var err error
		existent_files := []file.File{}
		if existent_files, err = fs.ReadDir(ctx, dest); err != nil {
			return nil, err
		}
		files = lo.Filter[file.File](files, func(item file.File, index int) bool {
			file_path, _ := item.Abs(ctx)
			return !lo.ContainsBy[file.File](existent_files, func(other file.File) bool {
				other_path, _ := other.Abs(ctx)
				return file_path == other_path
			})
		})
	}

	if opts.DryRun {
		return files, nil
	}

	if opts.Clean {
		backup_dir, err := os.MkdirTemp(os.TempDir(), "magus-*")
		defer fs.CleanDir(ctx, backup_dir)
		if err != nil {
			return nil, err
		}
		if err := fs.CopyDir(ctx, dest, backup_dir); err != nil {
			return nil, err
		}
		if err := fs.CleanDir(ctx, dest); err != nil {
			defer fs.CopyDir(ctx, backup_dir, dest)
			return nil, err
		}
	}

	if err := fs.WriteFiles(ctx, files); err != nil {
		return nil, err
	}
	return files, nil
}
