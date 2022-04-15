package imgge

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

func NewColorShift(r image.Rectangle, maxHeight, maxShift, n int) *ColorShift {
	return &ColorShift{
		Rect:      r,
		MaxHeight: maxHeight,
		MaxShift:  maxShift,
		N:         n,
		blocks:    randomColorShiftBlocks(r, maxHeight, maxShift, n),
	}
}

type ColorShift struct {
	Rect      image.Rectangle
	MaxHeight int
	MaxShift  int
	N         int
	blocks    []colorShiftBlock
}

func (c *ColorShift) Apply(img draw.Image) {
	src := img
	w, h := c.Rect.Max.X, c.Rect.Max.Y

	for _, block := range c.blocks {
		if block.isBlueShift {
			for x := w; x > block.shift; x-- {
				for y := block.y; y < block.y+block.rowHeight; y++ {
					if x-block.shift >= w || y > h {
						continue
					}

					_, _, b, _ := src.At(x-block.shift, y).RGBA()
					r0, g0, _, a0 := src.At(x, y).RGBA()
					img.Set(x, y, color.RGBA{uint8(r0), uint8(g0), uint8(b), uint8(a0)})
				}
			}
		} else {
			for x := block.shift; x < w; x++ {
				for y := block.y; y < block.y+block.rowHeight; y++ {
					if x+block.shift >= w || y > h {
						continue
					}

					r, _, _, _ := src.At(x+block.shift, y).RGBA()
					_, g0, b0, a0 := src.At(x, y).RGBA()
					img.Set(x, y, color.RGBA{uint8(r), uint8(g0), uint8(b0), uint8(a0)})
				}
			}
		}
	}
}

func (c *ColorShift) ApplyNext(img draw.Image) {}

func (c *ColorShift) Randomize() {
	c.blocks = randomColorShiftBlocks(c.Rect, c.MaxHeight, c.MaxShift, c.N)
}

func (c *ColorShift) Name() string { return "ColorShift" }

func randomColorShiftBlocks(r image.Rectangle, maxHeight, maxShift, n int) []colorShiftBlock {
	blocks := make([]colorShiftBlock, n)

	for i := range blocks {
		blocks[i] = colorShiftBlock{
			shift:       rand.Intn(maxShift),
			y:           rand.Intn(r.Dy()) - r.Min.Y,
			rowHeight:   rand.Intn(maxHeight),
			isBlueShift: bool(rand.Intn(2) == 1),
		}
	}

	return blocks
}

type colorShiftBlock struct {
	shift       int
	y           int
	rowHeight   int
	isBlueShift bool
}
