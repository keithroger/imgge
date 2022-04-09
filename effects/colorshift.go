package effects

import (
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

type ColorShift struct {
	imgWidth, imgHeight int
	MaxHeight           int
	MaxShift            int
	n                   int
	blocks              []colorShiftBlock
}

func NewColorShift(img draw.Image, maxHeight, maxShift, n int) ColorShift {
	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	blocks := make([]colorShiftBlock, n)

	rand.Seed(time.Now().UnixNano())

	for i := range blocks {
		blocks[i] = colorShiftBlock{
			shift:       rand.Intn(maxShift),
			y:           rand.Intn(imgHeight),
			rowHeight:   rand.Intn(maxHeight),
			isBlueShift: bool(rand.Intn(2) == 1),
		}
	}

	return ColorShift{
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
        MaxHeight: maxHeight,
		MaxShift:  maxShift,
		n:         n,
		blocks:    blocks,
	}
}

func (c *ColorShift) Apply(img draw.Image) {
	src := img

	for _, block := range c.blocks {
		if block.isBlueShift {
			for x := c.imgWidth; x > block.shift; x-- {
				for y := block.y; y < block.y+block.rowHeight; y++ {
					_, _, b, _ := src.At(x-block.shift, y).RGBA()
					r0, g0, _, a0 := src.At(x, y).RGBA()
					img.Set(x, y, color.RGBA{uint8(r0), uint8(g0), uint8(b), uint8(a0)})
				}
			}
		} else {
			for x := block.shift; x < c.imgWidth; x++ {
				for y := block.y; y < block.y+block.rowHeight; y++ {
					r, _, _, _ := src.At(x+block.shift, y).RGBA()
					_, g0, b0, a0 := src.At(x, y).RGBA()
					img.Set(x, y, color.RGBA{uint8(r), uint8(g0), uint8(b0), uint8(a0)})
				}
			}
		}
	}
}

type colorShiftBlock struct {
	shift       int
	y           int
	rowHeight   int
	isBlueShift bool
}
