package imgge

import (
	"image"
	"image/draw"
	"math/rand"
)

// PixelPop represents an effect that makes random pixels and draws a square of the same color.
type PixelPop struct {
	Rect             image.Rectangle
	MinSize, MaxSize int
	N                int
	blocks           []pixelPopBlock
}

// NewPixelPop returns a PixelPop struct.
func NewPixelPop(r image.Rectangle, minSize, maxSize, n int) *PixelPop {
	return &PixelPop{
		Rect:    r,
		MinSize: minSize,
		MaxSize: maxSize,
		N:       n,
		blocks:  randomPixelPopBlocks(r, minSize, maxSize, n),
	}
}

// Apply selects random pixels in the image and draws them as squares.
func (p *PixelPop) Apply(img draw.Image) {
	for _, block := range p.blocks {
		c := img.At(
			block.rect.Min.X+block.rect.Dx()/2,
			block.rect.Min.Y+block.rect.Dx()/2,
		)

		draw.Draw(img, block.rect, &image.Uniform{c}, image.Point{0, 0}, draw.Src)
	}
}

// Next makes small random changes to the position of the source pixels.
func (p *PixelPop) Next() {
	for i := range p.blocks {
		p.blocks[i].rect.Min.X += rand.Intn(3) - 1
		p.blocks[i].rect.Min.Y += rand.Intn(3) - 1
		p.blocks[i].rect.Max.X += rand.Intn(3) - 1
		p.blocks[i].rect.Max.Y += rand.Intn(3) - 1
	}
}

// Randomize reinitializes the positions of the shifted blocks.
func (p *PixelPop) Randomize() {
	p.blocks = randomPixelPopBlocks(p.Rect, p.MinSize, p.MaxSize, p.N)
}

// randomPixelPopBlocks is used to initialize positions of source pixels.
func randomPixelPopBlocks(r image.Rectangle, minSize, maxSize, n int) []pixelPopBlock {
	blocks := make([]pixelPopBlock, n)

	for i := range blocks {
		x, y := rand.Intn(r.Max.X), rand.Intn(r.Max.Y)
		squareWidth := rand.Intn(maxSize) - minSize
		rect := image.Rectangle{
			image.Point{x - squareWidth/2, y - squareWidth/2},
			image.Point{x + squareWidth/2, y + squareWidth/2},
		}

		blocks[i] = pixelPopBlock{
			squareWidth: squareWidth,
			rect:        rect,
		}
	}

	return blocks
}

// pixelPopBlock contains the source rectange rect with a size squareWidth.
type pixelPopBlock struct {
	squareWidth int
	rect        image.Rectangle
}
