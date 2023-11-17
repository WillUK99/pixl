package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colourPicker := SetupColourPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colourPicker)

	app.PixlWindow.SetContent(appLayout)
}
