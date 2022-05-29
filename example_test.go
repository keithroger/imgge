package imgge_test

import (
	"log"

	"github.com/keithroger/imgge"
)

func ExampleShift() {
	// Import jpeg using included function
	img, err := imgge.JpegToImage("images/original.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Create and Apply Effect
	effect := imgge.NewShift(img.Bounds(), 20, 30, 22)
	effect.Apply(img)

	// Export to png using included function
	err = imgge.SaveAsPng("Shift.png", img)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleColorShift() {
	// Import jpeg using included function
	img, err := imgge.JpegToImage("images/original.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Create and Apply Effect
	effect := imgge.NewColorShift(img.Bounds(), 20, 30, 25)
	effect.Apply(img)

	// Export to png using included function
	err = imgge.SaveAsPng("ColorShift.png", img)
	if err != nil {
		log.Fatal(err)
	}
}

func ExamplePixelSort() {
	// Import jpeg using included function
	img, err := imgge.JpegToImage("images/original.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Create and Apply Effect
	effect := imgge.NewPixelSort(img.Bounds(), 50, 100, "horiz")
	effect.Apply(img)

	// Export to png using included function
	err = imgge.SaveAsPng("PixelSort.png", img)
	if err != nil {
		log.Fatal(err)
	}
}

func ExamplePixelPop() {
	// Import jpeg using included function
	img, err := imgge.JpegToImage("images/original.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Create and Apply Effect
	effect := imgge.NewPixelPop(img.Bounds(), 15, 50, 100)
	effect.Apply(img)

	// Export to png using included function
	err = imgge.SaveAsPng("PixelPop.png", img)
	if err != nil {
		log.Fatal(err)
	}
}
