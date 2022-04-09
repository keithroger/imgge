package effects

import (
	"image/color"
	"image/draw"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

type PixelSort struct {
	imgWidth, imgHeight int
	maxLen              int
	n                   int
	orientation         string
	blocks              []pixelSortBlock
}

// Pixel sort rows are applied to image at current settings
func (p *PixelSort) Apply(img draw.Image) {
	for _, block := range p.blocks {
		pixels := block.getPixels(img, p.orientation)
		sort.Sort(ByLum(pixels))

		p.drawPxSort(img, block, pixels)
	}
}

func (p *PixelSort) drawPxSort(img draw.Image, block pixelSortBlock, pixels []color.Color) {
    for i, c := range pixels {
		if p.orientation == "horiz" {
			img.Set(block.x0+i, block.y0, c)
		} else if p.orientation == "vert" {
			img.Set(block.x0, block.y0+i, c)
		} else {
			log.Fatal("Orientation must be \"horiz\" or \"vert\"")
		}
	}
}

func NewPixelSort(img draw.Image, maxLen, n int, orientation string) PixelSort {
	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	rand.Seed(time.Now().UnixNano())

	blocks := make([]pixelSortBlock, n)

	for i := range blocks {
        x0 := rand.Intn(imgWidth)
        y0 := rand.Intn(imgHeight)
        length := rand.Intn(maxLen)


		blocks[i] = pixelSortBlock{
			x0:     x0,
			y0:     y0,
			length: length,
		}
	}

	return PixelSort{
		imgWidth:    imgWidth,
		imgHeight:   imgHeight,
		maxLen:      maxLen,
		n:           n,
		orientation: orientation,
		blocks:      blocks,
	}
}

type pixelSortBlock struct {
	x0, y0 int
	length int
}

func (p *pixelSortBlock) getPixels(img draw.Image, orientation string) []color.Color {
	pixels := make([]color.Color, p.length)

	for i := p.length - 1; i >= 0; i-- {
		if orientation == "horiz" {
			pixels[i] = img.At(p.x0+i, p.y0)
		} else if orientation == "vert" {
			pixels[i] = img.At(p.x0, p.y0+i)
		} else {
			log.Fatal("Orientation must be \"horiz\" or \"vert\"")
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
