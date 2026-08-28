package filemanager

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Project struct {
	Dirs  []Directory
	Files []File
}

type Service struct {
}

func (s *Service) Open(ctx context.Context, path string) (File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		slog.ErrorContext(ctx, "error occured while opening file", slog.String("error", err.Error()))
		return nil, fmt.Errorf("opening file: %w", err)
	}

	// Discuss how we should handle missing methods on files.
	return file, nil
}

func (s *Service) OpenFolder(ctx context.Context, root string) (*Project, error) {
	var project Project
	
	// This implementation is wrong.
	// Open folder should just open folder. But fileexplorer will handle walking
	// So we should open folder and return pointer to os.File with opened directory. 
	// Discuss it 
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			file, err := os.OpenFile(path, os.O_RDWR, 0)
			if err != nil {
				slog.ErrorContext(ctx, "error occured while opening directory", slog.String("error", err.Error()))
				return fmt.Errorf("opening folder: %w", err)
			}
			project.Dirs = append(project.Dirs, file)
		}

		return nil
	})

	if err != nil {
		if err != nil {
			slog.ErrorContext(ctx, "error occured while opening folder", slog.String("error", err.Error()))
			return nil, fmt.Errorf("opening folder: %w", err)
		}
	}

	return nil, nil
}

func New() *Service {
	return &Service{}
}
