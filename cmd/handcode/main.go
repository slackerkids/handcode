package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const logo = `
$$\   $$\                           $$\  $$$$$$\                  $$\           
$$ |  $$ |                          $$ |$$  __$$\                 $$ |          
$$ |  $$ | $$$$$$\  $$$$$$$\   $$$$$$$ |$$ /  \__| $$$$$$\   $$$$$$$ | $$$$$$\  
$$$$$$$$ | \____$$\ $$  __$$\ $$  __$$ |$$ |      $$  __$$\ $$  __$$ |$$  __$$\ 
$$  __$$ | $$$$$$$ |$$ |  $$ |$$ /  $$ |$$ |      $$ /  $$ |$$ /  $$ |$$$$$$$$ |
$$ |  $$ |$$  __$$ |$$ |  $$ |$$ |  $$ |$$ |  $$\ $$ |  $$ |$$ |  $$ |$$   ____|
$$ |  $$ |\$$$$$$$ |$$ |  $$ |\$$$$$$$ |\$$$$$$  |\$$$$$$  |\$$$$$$$ |\$$$$$$$\ 
\__|  \__| \_______|\__|  \__| \_______| \______/  \______/  \_______| \_______|    
                                                                                `

func main() {
	a := app.New()

	if len(os.Args) < 2 {
		log.Println("Provide filename")
		return
	}

	filename := os.Args[1]
	file, err := os.OpenFile(filename, os.O_RDWR, 0)
	if err != nil {
		log.Println("error opening file")
		return
	}

	text, err := io.ReadAll(file)
	if err != nil {
		log.Println("error reading file")
		return
	}

	file.Close()

	textArea := widget.NewMultiLineEntry()
	textArea.TextStyle = fyne.TextStyle{Monospace: true}
	textArea.SetText(string(text))

	w := a.NewWindow("Handcode")
	w.Resize(fyne.NewSize(1000, 500))

	saveFile := func() {
		err := os.WriteFile(filename, []byte(textArea.Text), 0644)
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
}
