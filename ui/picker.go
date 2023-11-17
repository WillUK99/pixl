package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
)

func SetupColourPicker(app *AppInit) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(c color.Color) {
		app.State.BrushColour = c
		app.Swatches[app.State.SwatchSelected].SetColour(c)
	})

	return container.NewVBox(picker)
}
