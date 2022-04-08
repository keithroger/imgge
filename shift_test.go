package imgge

import (
	"strconv"
	"testing"
)

func TestShift(t *testing.T) {
	for i := 1; i <= 4; i++ {
		img, err := openTestImage()
		if err != nil {
			t.Error(err)
		}

		shift := NewShift(img, 5*i, 25*i, 20)
		shift.Apply(img)

		err = outputTestImage(img, "TestShift"+strconv.Itoa(i)+".png")
		if err != nil {
			t.Error(err)
		}
	}
}
