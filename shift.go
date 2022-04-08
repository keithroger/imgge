package imgge

import (
	"image"
	"image/draw"
	"math/rand"
	"time"
)

type Shift struct {
	imgWidth, imgHeight int
	MaxHeight           int
	MaxShift            int
	n                   int
	blocks              []shiftBlock
}

func (s *Shift) Apply(img draw.Image) {
	src := img

	for _, block := range s.blocks {
		draw.Draw(img, block.rectangle1, src, block.srcPoint1, draw.Src)
		draw.Draw(img, block.rectangle2, src, block.srcPoint2, draw.Src)
	}
}

// Draws next frame to create a jiggling animation of rows.
func (s *Shift) ApplyNext(img draw.Image) {}

func NewShift(img draw.Image, maxHeight, maxShift, n int) Shift {
	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	blocks := make([]shiftBlock, n)

	rand.Seed(time.Now().UnixNano())

	for i := range blocks {
		randX := rand.Intn(maxShift)
		randY := rand.Intn(imgHeight)
		rowHeight := rand.Intn(maxHeight)

		// shift left or right randomly
		if rand.Intn(2) == 1 {
			blocks[i].srcPoint1 = image.Point{0, randY}
			blocks[i].rectangle1 = image.Rectangle{
				image.Point{randX, randY},
				image.Point{imgWidth, randY + rowHeight},
			}

			blocks[i].srcPoint2 = image.Point{imgWidth - randX, randY}
			blocks[i].rectangle2 = image.Rectangle{
				image.Point{0, randY},
				image.Point{randX, randY + rowHeight},
			}
		} else {
			blocks[i].srcPoint1 = image.Point{randX, randY}
			blocks[i].rectangle1 = image.Rectangle{
				image.Point{0, randY},
				image.Point{imgWidth - randX, randY + rowHeight},
			}

			blocks[i].srcPoint2 = image.Point{0, randY}
			blocks[i].rectangle2 = image.Rectangle{
				image.Point{imgWidth - randX, randY},
				image.Point{imgWidth, randY + rowHeight},
			}
		}
	}

	return Shift{
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
		MaxHeight: maxHeight,
		MaxShift:  maxShift,
		n:         n,
		blocks:    blocks,
	}
}

type shiftBlock struct {
	srcPoint1, srcPoint2   image.Point
	rectangle1, rectangle2 image.Rectangle
}
