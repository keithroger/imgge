package imgge

import (
	"strconv"
	"testing"
)

func TestPixelSort(t *testing.T) {
	for i := 1; i <= 4; i++ {
		img, err := openTestImage()
		if err != nil {
			t.Error(err)
		}

        effect := NewPixelSort(img, 90*i, 1000*i, "horiz")
        effect.Apply(img)

		err = outputTestImage(img, "TestPixelSort"+strconv.Itoa(i)+".png")
		if err != nil {
			t.Error(err)
		}
	}
}
