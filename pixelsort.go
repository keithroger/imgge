package imgge

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"math/rand"
	"sort"
)

// Pixelsort represents an Effect that moves sorts lines of pixels.
type PixelSort struct {
	Rect        image.Rectangle
	MaxLen      int
	N           int
	Orientation string
	blocks      []pixelSortBlock
}

// NewPixelSort returns a new PixelSort Effect.
func NewPixelSort(r image.Rectangle, maxLen, n int, orientation string) *PixelSort {
	return &PixelSort{
		Rect:        r,
		MaxLen:      maxLen,
		N:           n,
		Orientation: orientation,
		blocks:      randomPixelSortBlocks(r, maxLen, n, orientation),
	}
}

// Apply sorts pixels according to the PixelSort settings.
func (p *PixelSort) Apply(img draw.Image) {
	for _, block := range p.blocks {
		pixels := block.getPixels(img, p.Orientation)
		sort.Sort(byLum(pixels))

		p.drawPxSort(img, block, pixels)
	}
}

// Next makes small random changes to the source point for the PixelSort
func (p *PixelSort) Next() {
	for i := range p.blocks {
		p.blocks[i].srcPt.X += rand.Intn(3) - 1
		p.blocks[i].srcPt.Y += rand.Intn(3) - 1
	}
}

// Randomize reinitializes the positions of the shifted sections.
func (p *PixelSort) Randomize() {
	p.blocks = randomPixelSortBlocks(p.Rect, p.MaxLen, p.N, p.Orientation)
}

// randomPixelSortBlocks is used to initialize positions of the sorted areas.
func randomPixelSortBlocks(r image.Rectangle, maxLen, n int, orientation string) []pixelSortBlock {
	blocks := make([]pixelSortBlock, n)

	switch orientation {
	case "hoirz":
		if maxLen > r.Max.X {
			log.Fatal("Maxlen > image width")
		}
	case "vert":
		if maxLen > r.Max.X {
			log.Fatal("MaxLen > image height")
		}
	}

	for i := range blocks {
		var x, y, length int

		for {
			x = rand.Intn(r.Max.X)
			y = rand.Intn(r.Max.Y)
			length = rand.Intn(maxLen)

			if doesFit(x, y, r.Max.X, r.Max.Y, length, orientation) {
				break
			}
		}

		blocks[i] = pixelSortBlock{
			srcPt:  image.Point{x, y},
			length: length,
		}
	}

	return blocks
}

// drawPxSort is used to color an array of pixels to the image img.
func (p *PixelSort) drawPxSort(img draw.Image, block pixelSortBlock, pixels []color.Color) {
	for i, c := range pixels {
		if p.Orientation == horizontal {
			img.Set(block.srcPt.X+i, block.srcPt.Y, c)
		} else if p.Orientation == vertical {
			img.Set(block.srcPt.X, block.srcPt.Y+i, c)
		}
	}
}

// doesFit checks if the pixel sort will fit in when drawn to the image.
func doesFit(x, y, w, h, length int, orientation string) bool {
	if orientation == horizontal && x+length > w {
		return false
	} else if orientation == vertical && y+length > h {
		return false
	}

	return true
}

// pixelSortBlock represents where a the effect will be drawn.
type pixelSortBlock struct {
	srcPt  image.Point
	length int
}

// getPixels returns an array of pixels to be sorted.
func (p *pixelSortBlock) getPixels(img draw.Image, orientation string) []color.Color {
	pixels := make([]color.Color, p.length)

	for i := p.length - 1; i >= 0; i-- {
		switch orientation {
		case horizontal:
			pixels[i] = img.At(p.srcPt.X+i, p.srcPt.Y)
		case vertical:
			pixels[i] = img.At(p.srcPt.X, p.srcPt.Y+i)
		default:
			log.Fatalf("Orientation must be \"%s\" or \"%s\"", horizontal, vertical)
		}
	}

	return pixels
}

type byLum []color.Color

func (l byLum) Len() int { return len(l) }

func (l byLum) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l byLum) Less(i, j int) bool { return luminosity(l[i]) < luminosity(l[j]) }

func luminosity(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	return math.Sqrt(0.299*math.Pow(float64(r), 2) + 0.587*math.Pow(float64(g), 2) + 0.114*math.Pow(float64(b), 2))
}
