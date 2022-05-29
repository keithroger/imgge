package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/keithroger/imgge"
)

const (
	originalImg = "original.jpg"
)

type namedEffect struct {
	filename string
	effect   imgge.Effect
}

func main() {
	rand.Seed(1618)

	infile, err := os.Open(originalImg)
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()

	jpg, err := jpeg.Decode(infile)
	r := jpg.Bounds()

	effects := []namedEffect{
		{"Shift.png", imgge.NewShift(r, 20, 30, 22)},
		{"ColorShift.png", imgge.NewColorShift(r, 20, 30, 22)},
		{"PixelSort.png", imgge.NewPixelSort(r, 150, 150, "horiz")},
		{"PixelPop.png", imgge.NewPixelPop(r, 5, 30, 200)},
	}

	wg := sync.WaitGroup{}
	for _, fx := range effects {
		img := image.NewRGBA(r)
		draw.Draw(img, r, jpg, r.Min, draw.Src)

		wg.Add(1)
		go generate(img, fx, &wg)
	}
	wg.Wait()
}

func generate(img draw.Image, fx namedEffect, wg *sync.WaitGroup) {
	defer wg.Done()

	fx.effect.Apply(img)

	outfile, err := os.Create(fx.filename)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
