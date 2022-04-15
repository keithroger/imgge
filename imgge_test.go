package imgge_test

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"github.com/keithroger/imgge"
)

const (
	imgW, imgH = 400, 400
	imgDir     = "test_images"
)

func sampleImg() draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	draw.Draw(img, img.Bounds(), image.White, image.Point{0, 0}, draw.Src)

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
	rand.Seed(1984)

	img := sampleImg()

	tt := []imgge.Effect{
		imgge.NewShift(img, 20, 30, 25),
		imgge.NewColorShift(img, 20, 30, 25),
		imgge.NewPixelSort(img, 50, 100, "horiz"),
		imgge.NewPixelSort(img, 50, 100, "vert"),
		imgge.NewPixelPop(img, 15, 50, 50),
	}

	// test Apply function
	for i := range tt {
		name := reflect.TypeOf(tt[i]).Elem().Name()

		fmt.Printf("Testing %s\n", name)

		img := sampleImg()
		tt[i].Apply(img)
		outputPNG(img, strconv.Itoa(i)+name+"_Apply.png")
	}
}
