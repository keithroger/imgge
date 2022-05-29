// Package applies randomized glitch effects to images.
// Effects can be applied to individual images or applied in sequences for animatons.
package imgge

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
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

func JpegToImage(filename string) (draw.Image, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	jpg, err := jpeg.Decode(infile)
	r := jpg.Bounds()
	img := image.NewRGBA(r)
	draw.Draw(img, r, jpg, r.Min, draw.Src)

	return img, nil
}

func SaveAsPng(filename string, img image.Image) error {
	outfile, err := os.Create("example.png")
	if err != nil {
		return err
	}

	err = png.Encode(outfile, img)
	if err != nil {
		return err
	}

	return nil
}
