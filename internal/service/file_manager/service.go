package filemanager

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Project struct {
	Dirs  []Directory
	Files []File
}

type Service struct {
}

type DiskFile struct {
	path string
	*os.File
}

func (df *DiskFile) Path() string {
	return df.path
}

func (df *DiskFile) Rename(newPath string) error {
	if err := os.Rename(df.path, newPath); err != nil {
		slog.Error("error occured while opening file", slog.String("error", err.Error()))
		return fmt.Errorf("renaming file: %w", err)
	}

	return nil
}

func (df *DiskFile) Delete() error {
	return os.Remove(df.path)
}

func (s *Service) Open(path string) (File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		slog.Error("error occured while opening file", slog.String("error", err.Error()))
		return nil, fmt.Errorf("opening file: %w", err)
	}

	df := &DiskFile{
		path,
		file,
	}

	return df, nil
}

func (s *Service) OpenFolder(root string) (*Project, error) {
	var project Project

	// This implementation is wrong.
	// Open folder should just open folder. But fileexplorer will handle walking
	// So we should open folder and return pointer to os.File with opened directory.
	// Discuss it
	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	if info.IsDir() {
	// 		file, err := os.OpenFile(path, os.O_RDWR, 0)
	// 		if err != nil {
	// 			slog.Error(ctx, "error occured while opening directory", slog.String("error", err.Error()))
	// 			return fmt.Errorf("opening folder: %w", err)
	// 		}
	// 		project.Dirs = append(project.Dirs, file)
	// 	}

	// 	return nil
	// })

	// if err != nil {
	// 	return nil, err
	// }

	return &project, nil
}

func (s *Service) Save(file string, buffer io.Reader) error {
	osFile, err := os.Create(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(osFile, buffer)
	return err
}

func New() *Service {
	return &Service{}
}
