package domain

import (
	"os"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
	"github.com/ismtabo/magus/filetree"
	"github.com/ismtabo/magus/fs"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/manifest"
	"github.com/ismtabo/magus/variable"
	"github.com/samber/lo"
)

type GenerateOptions struct {
	DryRun    bool
	Overwrite bool
	Clean     bool
	Variables variable.Variables
}

func Generate(ctx context.Context, dest string, mfst manifest.Manifest, opts GenerateOptions) ([]file.File, error) {
	if dest == "" {
		dest = mfst.Root
	}

	mgc := manifest.NewMagic(ctx, mfst)

	files, err := mgc.Render(ctx.WithCwd(dest), magic.MagicRenderOptions{
		Variables: opts.Variables,
	})
	if err != nil {
		return nil, err
	}

	if err := filetree.AssertNotHaveWriteConflicts(ctx, files); err != nil {
		return nil, err
	}

	if !opts.Overwrite && !opts.Clean {
		existent_files, err := fs.ReadFiles(ctx.WithCwd(dest), files, fs.ReadFilesOptions{
			NoFailOnMissing: true,
		})
		if err != nil {
			return nil, err
		}
		files = lo.Filter(files, func(item file.File, index int) bool {
			file_path, _ := item.Abs(ctx)
			return !lo.ContainsBy(existent_files, func(other file.File) bool {
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
