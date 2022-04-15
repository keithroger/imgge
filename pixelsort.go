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

func NewPixelSort(r image.Rectangle, maxLen, n int, orientation string) *PixelSort {
	return &PixelSort{
		Rect:        r,
		MaxLen:      maxLen,
		N:           n,
		Orientation: orientation,
		blocks:      randomPixelSortBlocks(r, maxLen, n, orientation),
	}
}

type PixelSort struct {
	Rect        image.Rectangle
	MaxLen      int
	N           int
	Orientation string
	blocks      []pixelSortBlock
}

// Pixel sort rows are applied to image at current settings.
func (p *PixelSort) Apply(img draw.Image) {
	for _, block := range p.blocks {
		pixels := block.getPixels(img, p.Orientation)
		sort.Sort(ByLum(pixels))

		p.drawPxSort(img, block, pixels)
	}
}

func (p *PixelSort) ApplyNext(img draw.Image) {}

func (p *PixelSort) Randomize() {
	p.blocks = randomPixelSortBlocks(p.Rect, p.MaxLen, p.N, p.Orientation)
}

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
		var x0, y0, length int

		for {
			x0 = rand.Intn(r.Max.X)
			y0 = rand.Intn(r.Max.Y)
			length = rand.Intn(maxLen)

			if doesFit(x0, y0, r.Max.X, r.Max.Y, length, orientation) {
				break
			}
		}

		blocks[i] = pixelSortBlock{
			x0:     x0,
			y0:     y0,
			length: length,
		}
	}

	return blocks
}

func (p *PixelSort) drawPxSort(img draw.Image, block pixelSortBlock, pixels []color.Color) {
	for i, c := range pixels {
		if p.Orientation == horizontal {
			img.Set(block.x0+i, block.y0, c)
		} else if p.Orientation == vertical {
			img.Set(block.x0, block.y0+i, c)
		}
	}
}

func doesFit(x, y, w, h, length int, orientation string) bool {
	if orientation == horizontal && x+length > w {
		return false
	} else if orientation == vertical && y+length > h {
		return false
	}

	return true
}

type pixelSortBlock struct {
	x0, y0 int
	length int
}

func (p *pixelSortBlock) getPixels(img draw.Image, orientation string) []color.Color {
	pixels := make([]color.Color, p.length)

	for i := p.length - 1; i >= 0; i-- {
		switch orientation {
		case horizontal:
			pixels[i] = img.At(p.x0+i, p.y0)
		case vertical:
			pixels[i] = img.At(p.x0, p.y0+i)
		default:
			log.Fatalf("Orientation must be \"%s\" or \"%s\"", horizontal, vertical)
		}
	}

	return pixels
}

type ByLum []color.Color

func (l ByLum) Len() int { return len(l) }

func (l ByLum) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l ByLum) Less(i, j int) bool { return luminosity(l[i]) < luminosity(l[j]) }

func luminosity(c color.Color) float64 {
	r, g, b, _ := c.RGBA()

	return math.Sqrt(0.299*math.Pow(float64(r), 2) + 0.587*math.Pow(float64(g), 2) + 0.114*math.Pow(float64(b), 2))
}
