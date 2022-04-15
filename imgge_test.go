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
	"reflect"
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

	tt := []imgge.Effect{
		imgge.NewShift(r, 20, 30, 25),
		imgge.NewColorShift(r, 20, 30, 25),
		imgge.NewPixelSort(r, 50, 100, "horiz"),
		imgge.NewPixelSort(r, 50, 100, "vert"),
		imgge.NewPixelPop(r, 15, 50, 50),
	}

	for _, tc := range tt {
		effect := tc

		effectName := reflect.TypeOf(effect).Elem().Name()

		// Test Apply() from Effects interface.
		tName := effectName + "Apply"
		t.Run(effectName, func(t *testing.T) {
			t.Parallel()

			img := sampleImg(r)
			effect.Apply(img)
			outputPNG(img, tName+"Apply.png")

		// Test Randomize() from Effects interface.
			img = sampleImg(r)
			effect.Randomize()
			effect.Apply(img)
			outputPNG(img, effectName +"Randomize.png")
		})
	}
}
