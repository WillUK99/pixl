package swatch

import (
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	Selected    bool
	Colour      color.Color
	SwatchIndex int
	handleClick func(s *Swatch)
}

func (s *Swatch) SetColour(c color.Color) {
	s.Colour = c
	s.Refresh() // part of the fyne package to redraw
}

func NewSwatch(state *apptype.State, colour color.Color, swatchIndex int, handleClick func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:    false,
		Colour:      colour,
		SwatchIndex: swatchIndex,
		handleClick: handleClick,
	}

	swatch.ExtendBaseWidget(swatch) // supplies fyne with the state required to setup a widget

	return swatch
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Colour)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  s,
	}
}
