package effects

import (
	"image"
	"image/draw"
	"math/rand"
    "time"
)

type PixelPop struct {
    imgWidth, imgHeight int
    minSize, maxSize int
    n int
    blocks []pixelPopBlock
}

func (p *PixelPop) Apply(img draw.Image) {
    for _, block := range p.blocks {
        c := img.At(block.x, block.y)

        draw.Draw(img, block.rect, &image.Uniform{c}, image.Point{}, draw.Src)
    }
}

func (p *PixelPop) ApplyNext(img draw.Image) {}

func (p *PixelPop) Randomize() {}

func (p *PixelPop) Name() string { return "pixelpop"}

func NewPixelPop(img draw.Image, minSize, maxSize, n int) *PixelPop {
	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	rand.Seed(time.Now().UnixNano())

	blocks := make([]pixelPopBlock, n)

    for i := range blocks {
        x, y := rand.Intn(imgWidth), rand.Intn(imgHeight)
        squareWidth := rand.Intn(maxSize) - minSize
		rect := image.Rectangle{
			image.Point{x - squareWidth/2, y - squareWidth/2},
			image.Point{x + squareWidth/2, y + squareWidth/2},
		}

        blocks[i] = pixelPopBlock{
            x: x,
            y: y,
            squareWidth: squareWidth,
            rect: rect,
        }
    }


    return &PixelPop{
        imgWidth: imgWidth,
        imgHeight: imgHeight,
        minSize: minSize,
        maxSize: maxSize,
        n: n,
        blocks: blocks,
    }
}

/*
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

	}
}
*/

type pixelPopBlock struct {
    x, y int
    squareWidth int
    rect image.Rectangle
}
