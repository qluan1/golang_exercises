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
	dy := 0.5/height*(ymax - ymin)
	dx := 0.5/width*(xmax-xmin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z
			img.Set(px, py, averageColor(
				mandelbrot(complex(x + dx, y + dy)),
				mandelbrot(complex(x + dx, y - dy)),
				mandelbrot(complex(x - dx, y + dy)),
				mandelbrot(complex(x - dx, y - dy)),
			))
		}
	}

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if  cmplx.Abs(v) > 2 {
			return getColor(int(n))
		}
	}
	return color.RGBA{0, 0, 0, 255}
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

func averageColor(c1, c2, c3, c4 color.RGBA) color.RGBA {
	r := uint32(c1.R) + uint32(c2.R) + uint32(c3.R) + uint32(c4.R)
	g := uint32(c1.G) + uint32(c2.G) + uint32(c3.G) + uint32(c4.G)
	b := uint32(c1.B) + uint32(c2.B) + uint32(c3.B) + uint32(c4.B)
	a := uint32(c1.A) + uint32(c2.A) + uint32(c3.A) + uint32(c4.A)

	return color.RGBA{
		R: uint8(r / 4),
		G: uint8(g / 4),
		B: uint8(b / 4),
		A: uint8(a / 4),
	}
}