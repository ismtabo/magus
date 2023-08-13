package imports

import "github.com/pkg/errors"

type ImportError interface {
	error
}

func AlreadyImportedError(i Import) ImportError {
	return errors.Errorf("already imported %s", i.To().Path())
}
