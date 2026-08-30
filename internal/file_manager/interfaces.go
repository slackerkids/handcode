package filemanager

import (
	"io"
)

type FileManager interface {
	Open(path string) (*File, error)
	OpenFolder(root string) (*Project, error)
	Save(file string, buffer io.Reader) error
	SaveAll(files []File) error
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
	Rename(newPath string) error
	io.ReadWriteCloser
}
