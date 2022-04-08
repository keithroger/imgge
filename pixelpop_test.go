package imgge

import (
	"strconv"
	"testing"
)

func TestPixelPop(t *testing.T) {
	for i := 1; i <= 4; i++ {
		img, err := openTestImage()
		if err != nil {
			t.Error(err)
		}

		PixelPop(img, i, 10*i, 1000*i)

		err = outputTestImage(img, "TestPixelPop"+strconv.Itoa(i)+".png")
		if err != nil {
			t.Error(err)
		}
	}
}
