package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if  cmplx.Abs(v) > 2 {
			return getColor(int(n))
		}
	}
	return color.Black
}

func getColor(n int) color.RGBA {
	if n == 2 {
		return color.RGBA{0, 0, 0, 255} // Black
	}

	segment := 198 / 4 // Divided by number of colors-1 (5 colors - 1 = 4)
	switch {
	case n < 2+segment: // From black to blue
		return color.RGBA{uint8(255 * (n-2) / segment), 0, 255, 255}
	case n < 2+2*segment: // From blue to green
		return color.RGBA{0, uint8(255 * (n-2-segment) / segment), 255, 255}
	case n < 2+3*segment: // From green to yellow
		return color.RGBA{uint8(255 * (n-2-2*segment) / segment), 255, uint8(255 * (1 - float64(n-2-2*segment)/float64(segment))), 255}
	default: // From yellow to red
		return color.RGBA{255, uint8(255 * (1 - float64(n-2-3*segment)/float64(segment))), 0, 255}
	}
}
