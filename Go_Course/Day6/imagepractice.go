package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

func solidImage(width int, height int, r, g, b, a byte) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height*img.Stride; i += 4 {
		img.Pix[i] = r
		img.Pix[i+1] = g
		img.Pix[i+2] = b
		img.Pix[i+3] = a
	}
	return img

}

func drawSolidColorRect(img *image.NRGBA, bounds image.Rectangle, r, g, b, a byte) {
	height := bounds.Max.Y - bounds.Min.Y
	width := bounds.Max.X - bounds.Min.X
	newColor := solidImage(width, height, r, g, b, a)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, newColor.At(0, 0))
		}
	}
}

func apparentColor(color, bgColor, alpha, bgAlpha byte) byte {
	return byte(float64(color)*float64(alpha) + (1.0-float64(alpha))*(float64(bgColor)*float64(bgAlpha)))
}

func compositeColor(apparentColor, compositeAlpha byte) byte {
	endColor := float64(apparentColor) / float64(compositeAlpha)
	return byte(endColor)
}

func compositeAlpha(alpha, bgAlpha byte) byte {
	compositeAlpha := float64(alpha) + (1.0-float64(alpha))*float64(bgAlpha)
	return byte(compositeAlpha)
}

func AlphaBlend(top *image.NRGBA, bottom *image.NRGBA) {
	topBounds := top.Bounds()
	bottomBounds := bottom.Bounds()
	minX := int(math.Max(float64(topBounds.Min.X), float64(bottomBounds.Min.X)))
	fmt.Println(minX)
	minY := int(math.Max(float64(topBounds.Min.Y), float64(bottomBounds.Min.Y)))
	fmt.Println(minY)
	maxX := int(math.Min(float64(topBounds.Max.X), float64(bottomBounds.Max.X)))
	fmt.Println(maxX)
	maxY := int(math.Min(float64(topBounds.Max.Y), float64(bottomBounds.Max.Y)))
	fmt.Println(maxY)
	rect := image.Rect(minX, minY, maxX, maxY)
	img := image.NewRGBA(rect)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			topIdx := top.Stride*y + x*4
			bottomIdx := bottom.Stride*y + x*4
			idx := img.Stride*y + x*4
			// blending
			cA := compositeAlpha(top.Pix[topIdx+3], bottom.Pix[bottomIdx+3])
			img.Pix[idx+3] = cA
			for i := 0; i < 3; i++ {
				aC := apparentColor(top.Pix[topIdx+i],
					bottom.Pix[bottomIdx+i],
					top.Pix[topIdx+3],
					bottom.Pix[bottomIdx+3],
				)
				cC := aC // cA
				top.Pix[topIdx+i] = cC
			}
		}
	}

}

func main() {
	fid, err := os.OpenFile("test.png", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer fid.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	img := solidImage(1000, 500, 255, 0, 0, 128)
	img2 := solidImage(500, 150, 0, 120, 255, 255)

	AlphaBlend(img, img2)

	png.Encode(fid, img)
}
