package imgge

import (
	"image/color"
	"image/draw"
	"log"
	"math"
	"math/rand"
	"sort"
)

const (
	vertical   = "vert"
	horizontal = "horiz"
)

type PixelSort struct {
	imgWidth, imgHeight int
	maxLen              int
	n                   int
	orientation         string
	blocks              []pixelSortBlock
}

// Pixel sort rows are applied to image at current settings.
func (p *PixelSort) Apply(img draw.Image) {
	for _, block := range p.blocks {
		pixels := block.getPixels(img, p.orientation)
		sort.Sort(ByLum(pixels))

		p.drawPxSort(img, block, pixels)
	}
}

func (p *PixelSort) ApplyNext(img draw.Image) {}

func (p *PixelSort) Randomize() {}

func (p *PixelSort) drawPxSort(img draw.Image, block pixelSortBlock, pixels []color.Color) {
	for i, c := range pixels {
		if p.orientation == horizontal {
			img.Set(block.x0+i, block.y0, c)
		} else if p.orientation == vertical {
			img.Set(block.x0, block.y0+i, c)
		}
	}
}

func NewPixelSort(img draw.Image, maxLen, n int, orientation string) *PixelSort {
	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	blocks := make([]pixelSortBlock, n)

	switch orientation {
	case "hoirz":
		if maxLen > imgWidth {
			log.Fatal("Maxlen > image width")
		}
	case "vert":
		if maxLen > imgHeight {
			log.Fatal("MaxLen > image height")
		}
	}

	for i := range blocks {
		var x0, y0, length int

		for {
			x0 = rand.Intn(imgWidth)
			y0 = rand.Intn(imgHeight)
			length = rand.Intn(maxLen)

			if doesFit(x0, y0, imgWidth, imgHeight, length, orientation) {
				break
			}
		}

		blocks[i] = pixelSortBlock{
			x0:     x0,
			y0:     y0,
			length: length,
		}
	}

	return &PixelSort{
		imgWidth:    imgWidth,
		imgHeight:   imgHeight,
		maxLen:      maxLen,
		n:           n,
		orientation: orientation,
		blocks:      blocks,
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
