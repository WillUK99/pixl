package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (r *SwatchRenderer) MinSize() fyne.Size {
	return r.square.MinSize()
}

func (r *SwatchRenderer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *SwatchRenderer) Refresh() {
	r.Layout(fyne.NewSize(20, 20))
	r.square.FillColor = r.parent.Colour

	if r.parent.Selected {
		r.square.StrokeColor = color.NRGBA{255, 255, 255, 255}
		r.square.StrokeWidth = 2
		r.objects[0] = &r.square // set the swatch squares to updated squares
	}

	if !r.parent.Selected {
		r.square.StrokeWidth = 0
		r.objects[0] = &r.square // set the swatch squares to updated squares
	}

	canvas.Refresh(r.parent)
}

func (r *SwatchRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *SwatchRenderer) Destroy() {}
