package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	var wid widget.CodeBlockSegment

	wid.Text = logo
	canvas := wid.Visual()

	a := app.New()
	w := a.NewWindow("Handcode")

	w.SetContent(canvas)
	w.Resize(fyne.NewSize(1000, 500))

	w.ShowAndRun()
}
