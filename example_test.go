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
	err = imgge.SaveAsPng("example.png", img)
	if err != nil {
		log.Fatal(err)
	}
}
