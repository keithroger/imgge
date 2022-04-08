package imgge

import (
	"image/draw"
)

type Effect interface {
	// draws the effect to the image with the structs current settings
	Apply(draw.Image)

	// draws a new frame to make image animated
	ApplyNext(draw.Image)

	// resets random components of effect
	Randomize()
}
