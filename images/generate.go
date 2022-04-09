// Run all effects and create create output images to arugment location
package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/keithroger/imgge"
	"github.com/keithroger/imgge/effects"
)

var (
	effectSelected = flag.String("name", "all", "name of effect")
	fileName       = flag.String("filename", "original.jpg", "jpeg file to demonstrate effects")
)

func main() {
	infile, err := os.Open(*fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer infile.Close()

	jpg, err := jpeg.Decode(infile)
	if err != nil {
		log.Fatalln(err)
	}

    img := image.NewRGBA(jpg.Bounds())
    draw.Draw(img, img.Bounds(), jpg, image.Point{}, draw.Src)

	examples := []imgge.Effect{
		effects.NewShift(img, 20, 30, 25),
        effects.NewColorShift(img, 20, 30, 25),
        effects.NewPixelSort(img, 50, 3000, "horiz"),
        effects.NewPixelPop(img, 3, 20, 700),
	}

	if !contains(examples, effectSelected) {
		log.Fatalf("%s is not contained in example list.\n", *effectSelected)
	}

	for i := range examples {
		if *effectSelected == "all" || *effectSelected == examples[i].Name() {
            img := image.NewRGBA(jpg.Bounds())
            draw.Draw(img, img.Bounds(), jpg, image.Point{}, draw.Src)

			fmt.Printf("Creating %s example\n", examples[i].Name())
			examples[i].Apply(img)
			outfile, err := os.Create("example" + strconv.Itoa(i) + "_" + examples[i].Name() + ".png")
			if err != nil {
				log.Fatalln(err)
			}

			err = png.Encode(outfile, img)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func contains(examples []imgge.Effect, effectSelected *string) bool {
	if *effectSelected == "all" {
		return true
	}

	for i := range examples {
		if examples[i].Name() == *effectSelected {
			return true
		}
	}

	return false
}
