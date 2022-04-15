package imgge_test

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/keithroger/imgge"
)

const (
	imgW, imgH = 400, 400
	imgDir     = "test_images"
)

func sampleImg(r image.Rectangle) draw.Image {
	img := image.NewRGBA(r)
	draw.Draw(img, img.Bounds(), image.White, image.Point{0, 0}, draw.Src)

	// Fill image with a chessboard pattern.
	for i := 0; i < imgW; i++ {
		for j := 0; j < imgH; j++ {
			if (i%100 < 50 && j%100 < 50) || ((i+50)%100 < 50 && (j+50)%100 < 50) {
				img.Set(i, j, color.Black)
			}
		}
	}

	return img
}

func outputPNG(img image.Image, filename string) {
	outfile, err := os.Create(filepath.Join(imgDir, filename))
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func TestEffects(t *testing.T) {
	t.Parallel()
	rand.Seed(1984)

	r := image.Rect(0, 0, imgW, imgH)

	tt := []struct {
		name   string
		effect imgge.Effect
	}{
		{"Shift", imgge.NewShift(r, 20, 30, 25)},
		{"ColorShift", imgge.NewColorShift(r, 20, 30, 25)},
		{"PixelSortHoriz", imgge.NewPixelSort(r, 50, 100, "horiz")},
		{"PixelSortVert", imgge.NewPixelSort(r, 50, 100, "vert")},
		{"PixelPop", imgge.NewPixelPop(r, 15, 50, 100)},
	}

	for _, tc := range tt {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Test Apply() from Effects interface.
			img := sampleImg(r)
			tc.effect.Apply(img)
			outputPNG(img, tc.name+"Apply.png")

			// Test Next() from Effect interface.
			img = sampleImg(r)
			tc.effect.Next()
			tc.effect.Apply(img)
			outputPNG(img, tc.name+"Next.png")

			// Test Randomize() from Effect interface.
			img = sampleImg(r)
			tc.effect.Randomize()
			tc.effect.Apply(img)
			outputPNG(img, tc.name+"Randomize.png")
		})
	}
}
