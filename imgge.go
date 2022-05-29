// Package applies randomized glitch effects to images.
// Effects can be applied to individual images or applied in sequences for animatons.
package imgge

import (
	"image/draw"
)

const (
	vertical   = "vert"
	horizontal = "horiz"
)

// Effect provides methods for applying effects and randomization.
type Effect interface {
	// Draws the effect to the image with with the current settings defined by the struct.
	Apply(draw.Image)

	// Next makes small variations to the effect.
	// Use in a sequence of images to produce an animated effect.
	Next()

	// Resets random components of effect.
	Randomize()
}
