package filetree

import (
	"sort"
	"strings"

	go_errors "github.com/pkg/errors"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/file"
)

func AssertNotHaveWriteConficts(ctx context.Context, files []file.File) error {
	paths := sort.StringSlice(make([]string, len(files)))
	for i, f := range files {
		path, err := f.Abs(ctx)
		if err != nil {
			// TODO: wrap error
			return err
		}
		paths[i] = path
	}
	paths.Sort()
	for i, f := range paths {
		if i == len(paths)-1 {
			break
		}
		for j := i + 1; j < len(paths); j++ {
			if isWriteConflict(f, paths[j]) {
				// TODO: wrap error
				return go_errors.Errorf("write conflict: %s and %s", f, paths[j])
			}
		}
	}
	return nil
}

func isWriteConflict(path1, path2 string) bool {
	return strings.HasPrefix(path2, path1)
}
