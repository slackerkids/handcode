package filemanager

import (
	"context"
	"io"
)

type FileManager interface {
	Open(ctx context.Context, path string) (File, error)
	OpenFolder(ctx context.Context, root string) (*Project, error)
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
