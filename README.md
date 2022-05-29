# Imgge

This module creates randomized glitch effects that can be applied to images
or on consecutive frames for animations.

## Documentation
[pkg.go.dev/github.com/keithroger/imgge](https://pkg.go.dev/github.com/keithroger/imgge)

## Install
```
$ go get github.com/keithroger/imgge
```


## Effects
All effects impliment the Effect interface.
```
type Effect interface {
	// Draws the effect to the image with with the current
    // settings defined by the struct.
	Apply(draw.Image)

	// Next makes small variations to the effect.
	// Use in a sequence of images to produce an animated effect.
	Next()

	// Resets random components of effect.
	Randomize()
```

### Shift
![Shift Effect](images/Shift.png "Shift Effect")

<details>
<summary>View Code</summary>

```
func Example() {
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
```
</details>

<br>

### Color Shift
![Color Shift](images/ColorShift.png "Color Shift Effect")

### Pixel Sort
![Pixel Sort](images/PixelSort.png "Pixel Sort Effect")

### Pixel Pop
![Pixel Pop](images/PixelPop.png "Pixel Pop Effect")

