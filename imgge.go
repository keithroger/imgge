package imgge

import (
	"image/draw"
)

const (
	vertical   = "vert"
	horizontal = "horiz"
)

type Effect interface {
	// Draws the effect to the image with the structs current settings.
	Apply(draw.Image)

	// Draws a new frame to make image animated.
	Next()

	// Resets random components of effect.
	Randomize()
}
