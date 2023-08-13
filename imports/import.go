package imports

import "github.com/ismtabo/magus/file"

type Import interface {
	From() file.File
	To() file.File
}

type importImpl struct {
	from file.File
	to   file.File
}

func NewImport(from, to file.File) Import {
	return &importImpl{
		from: from,
		to:   to,
	}
}

func (i *importImpl) From() file.File {
	return i.from
}

func (i *importImpl) To() file.File {
	return i.to
}
