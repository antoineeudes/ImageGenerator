package main

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"math"
	graphics "github.com/antoineeudes/graphics-go/graphics"
)

func main() {
    width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	srcImgRectangle := image.Rectangle{upLeft, lowRight}

	srcImg := image.NewRGBA(srcImgRectangle)
	maxSize := int(math.Ceil(math.Max(float64(srcImgRectangle.Dy()), float64(srcImgRectangle.Dx())) * math.Sqrt(2)))
	dstImg := image.NewRGBA(image.Rect(0, 0, maxSize, maxSize))

	cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				srcImg.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				srcImg.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	graphics.Rotate(dstImg, srcImg, &graphics.RotateOptions{math.Pi / 4.0})

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, dstImg)
}