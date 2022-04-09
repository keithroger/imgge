package imgge

import (
	"image"
	"image/draw"
	"math/rand"
)

func PixelPop(img draw.Image, minSize, maxSize, n int) {
	bound := img.Bounds()
	width := bound.Max.X
	height := bound.Max.Y

	for i := 0; i < n; i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		squareWidth := rand.Intn(maxSize) - minSize
		c := img.At(x, y)

		rect := image.Rectangle{
			image.Point{x - squareWidth/2, y - squareWidth/2},
			image.Point{x + squareWidth/2, y + squareWidth/2},
		}

		draw.Draw(img, rect, &image.Uniform{c}, image.Point{}, draw.Src)
	}
}
