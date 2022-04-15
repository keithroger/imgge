package imgge

import (
	"image"
	"image/draw"
	"math/rand"
)

func NewPixelPop(r image.Rectangle, minSize, maxSize, n int) *PixelPop {
	return &PixelPop{
		Rect:    r,
		MinSize: minSize,
		MaxSize: maxSize,
		N:       n,
		blocks:  randomPixelPopBlocks(r, minSize, maxSize, n),
	}
}

type PixelPop struct {
	Rect             image.Rectangle
	MinSize, MaxSize int
	N                int
	blocks           []pixelPopBlock
}

func (p *PixelPop) Apply(img draw.Image) {
	for _, block := range p.blocks {
		c := img.At(block.x, block.y)

		draw.Draw(img, block.rect, &image.Uniform{c}, image.Point{0, 0}, draw.Src)
	}
}

func (p *PixelPop) ApplyNext(img draw.Image) {}

func (p *PixelPop) Randomize() {
    randomPixelPopBlocks(p.Rect, p.MinSize, p.MaxSize, p.N)
}

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
			x:           x,
			y:           y,
			squareWidth: squareWidth,
			rect:        rect,
		}
	}

	return blocks
}

type pixelPopBlock struct {
	x, y        int
	squareWidth int
	rect        image.Rectangle
}