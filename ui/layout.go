package ui

func Setup(app *AppInnit) {
	swatchesContainer := BuildSwatches(app)

	app.PixlWindow.SetContent(swatchesContainer)
}
