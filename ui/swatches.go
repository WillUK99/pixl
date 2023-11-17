package ui

import (
	"image/color"
	"pixl/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(canvasSwatches); i++ {
		initialColour := color.NRGBA{0, 0, 0, 255}
		s := swatch.NewSwatch(app.State, initialColour, i, func(s *swatch.Swatch) {
			// Remove any highlights around the swatches
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			// We need to select a new selected swatch and update the brush colour to be === the swatch colour
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColour = s.Colour
		})
		// By default the 1st swatch will be selected
		if i == 0 {
			s.Selected = true
			app.State.BrushColour = s.Colour
			s.Refresh()
		}
		// We update the swatches in the app.Swatches in order for us to later read and write. If we did not do this here then we would lose this data </3
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}
	// We can only use the fyne ui layouts with slices of CanvasObject
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
