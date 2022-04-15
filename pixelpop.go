package imgge

import (
	"image"
	"image/draw"
	"math/rand"
)

type PixelPop struct {
	imgWidth, imgHeight int
	minSize, maxSize    int
	n                   int
	blocks              []pixelPopBlock
}

func (p *PixelPop) Apply(img draw.Image) {
	for _, block := range p.blocks {
		c := img.At(block.x, block.y)

		draw.Draw(img, block.rect, &image.Uniform{c}, image.Point{0, 0}, draw.Src)
	}
}

func (p *PixelPop) ApplyNext(img draw.Image) {}

func (p *PixelPop) Randomize() {}

func NewPixelPop(img draw.Image, minSize, maxSize, n int) *PixelPop {
	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	blocks := make([]pixelPopBlock, n)

	for i := range blocks {
		x, y := rand.Intn(imgWidth), rand.Intn(imgHeight)
		squareWidth := rand.Intn(maxSize) - minSize
		rect := image.Rectangle{
			image.Point{x - squareWidth/2, y - squareWidth/2},
			image.Point{x + squareWidth/2, y + squareWidth/2},
		}

		blocks[i] = pixelPopBlock{
			x:           x,
			y:           y,
			squareWidth: squareWidth,
			rect:        rect,
		}
	}

	return &PixelPop{
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
		minSize:   minSize,
		maxSize:   maxSize,
		n:         n,
		blocks:    blocks,
	}
}

type pixelPopBlock struct {
	x, y        int
	squareWidth int
	rect        image.Rectangle
}
