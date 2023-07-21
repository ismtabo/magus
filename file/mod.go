package file

type File struct {
	// Path is the path of the file.
	Path string
	// Content is the content of the file.
	Value string
}

// NewFile creates a new file.
func NewFile(path, value string) File {
	return File{
		Path:  path,
		Value: value,
	}
}
