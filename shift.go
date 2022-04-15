package imgge

import (
	"image"
	"image/draw"
	"math/rand"
)

func NewShift(r image.Rectangle, maxHeight, maxShift, n int) *Shift {
	return &Shift{
		Rect:      r,
		MaxHeight: maxHeight,
		MaxShift:  maxShift,
		N:         n,
		blocks:    randomShiftBlocks(r, maxHeight, maxShift, n),
	}
}

type Shift struct {
	Rect      image.Rectangle
	MaxHeight int
	MaxShift  int
	N         int
	blocks    []shiftBlock
}

func (s *Shift) Apply(img draw.Image) {
	src := img

	for _, block := range s.blocks {
		draw.Draw(img, block.rectangle, src, block.srcPt, draw.Src)
	}
}

// Draws next frame to create a jiggling animation of rows.
func (s *Shift) Next() {
	for i := range s.blocks {
		s.blocks[i].srcPt.X += rand.Intn(3) - 1
		s.blocks[i].srcPt.Y += rand.Intn(3) - 1
	}
}

func (s *Shift) Randomize() {
	s.blocks = randomShiftBlocks(s.Rect, s.MaxHeight, s.MaxShift, s.N)
}

func randomShiftBlocks(r image.Rectangle, maxHeight, maxShift, n int) []shiftBlock {
	imgW := r.Max.X
	imgH := r.Max.Y

	blocks := make([]shiftBlock, n)

	for i := range blocks {
		randX := rand.Intn(maxShift)
		randY := rand.Intn(imgH)
		rowHeight := rand.Intn(maxHeight)

		// shift left or right randomly
		if rand.Intn(2) == 1 {
			blocks[i].srcPt = image.Point{0, randY}
			blocks[i].rectangle = image.Rectangle{
				image.Point{randX, randY},
				image.Point{imgW, randY + rowHeight},
			}
		} else {
			blocks[i].srcPt = image.Point{randX, randY}
			blocks[i].rectangle = image.Rectangle{
				image.Point{0, randY},
				image.Point{imgH - randX, randY + rowHeight},
			}
		}
	}

	return blocks
}

type shiftBlock struct {
	srcPt     image.Point
	rectangle image.Rectangle
}
