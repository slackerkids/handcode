package filemanager

import "io"

type FileManager interface {
	Open(path string) (File, error)
	OpenFolder(path string) (*Project, error)
	Save(File) error
	SaveAs(File) error
	SaveAll([]File) error
}

type Directory interface {
	Path() string
	Children() ([]Directory, []File)
	Delete() error
	Rename() error
}

type File interface {
	Path() string
	Delete() error
	Rename() error
	io.ReadWriteCloser
}

type Project struct {
	Dirs  []Directory
	Files []File
}
