package imgge

import (
	"strconv"
	"testing"
)

func TestColorShift(t *testing.T) {
	for i := 1; i <= 4; i++ {
		img, err := openTestImage()
		if err != nil {
			t.Error(err)
		}

        cShift := NewColorShift(img, 5*i, 25*i, 20)
        cShift.Apply(img)

		err = outputTestImage(img, "TestColorShift"+strconv.Itoa(i)+".png")
		if err != nil {
			t.Error(err)
		}
	}
}
