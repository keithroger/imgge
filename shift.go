package imgge

import (
	"image"
	"image/draw"
	"math/rand"
)

// Shift represents an effect that moves blocks within the images randomly left or right
type Shift struct {
	Rect      image.Rectangle
	MaxHeight int
	MaxShift  int
	N         int
	blocks    []shiftBlock
}

// NewShift creates a Shift struct.
// The Effect will be drawn within the rectangle r with shifted blocks
// max height of maxHeight and max horizontal shift of maxShift.
func NewShift(r image.Rectangle, maxHeight, maxShift, n int) *Shift {
	return &Shift{
		Rect:      r,
		MaxHeight: maxHeight,
		MaxShift:  maxShift,
		N:         n,
		blocks:    randomShiftBlocks(r, maxHeight, maxShift, n),
	}
}

// Apply shifts sections of the image according to the Shift settings.
func (s *Shift) Apply(img draw.Image) {
	src := img

	for _, block := range s.blocks {
		draw.Draw(img, block.rectangle, src, block.srcPt, draw.Src)
	}
}

// Next makes small random changes to the position of the shifted blocks.
func (s *Shift) Next() {
	for i := range s.blocks {
		s.blocks[i].srcPt.X += rand.Intn(3) - 1
		s.blocks[i].srcPt.Y += rand.Intn(3) - 1
	}
}

// Randomize reinitializes the positions of the shifted blocks.
func (s *Shift) Randomize() {
	s.blocks = randomShiftBlocks(s.Rect, s.MaxHeight, s.MaxShift, s.N)
}

// randomShiftBlocks is used to initialize position of shifted areas.
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

// shiftBlock contains the source and destination of the shifted block.
type shiftBlock struct {
	srcPt     image.Point
	rectangle image.Rectangle
}
