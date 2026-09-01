package app

import (
	"fmt"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	filemanager "github.com/slackerkids/handcode/internal/file_manager"
	"github.com/slackerkids/handcode/internal/validation"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	fyneapp := fyneapp.New()

	if err := validation.ValidateInput(os.Args); err != nil {
		return err
	}

	filename := os.Args[1]

	s := filemanager.New()
	file, err := s.Open(filename)

	text, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	textArea := widget.NewMultiLineEntry()
	textArea.TextStyle = fyne.TextStyle{Monospace: true}
	textArea.SetText(string(text))

	w := fyneapp.NewWindow("Handcode")
	w.Resize(fyne.NewSize(1000, 500))

	saveFile := func() {
		err := s.Save(filename, strings.NewReader(textArea.Text))
		if err != nil {
			fmt.Println("Error saving file:", err)
			return
		}
	}

	saveItem := fyne.NewMenuItem("Save", saveFile)

	saveItem.Shortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyS,
		Modifier: fyne.KeyModifierShortcutDefault,
	}

	fileMenu := fyne.NewMenu("File", saveItem)
	w.SetMainMenu(fyne.NewMainMenu(fileMenu))

	w.SetContent(textArea)
	w.ShowAndRun()

	return nil
}
