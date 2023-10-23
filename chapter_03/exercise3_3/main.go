package main

import (
	"fmt"
	"math"
	"image/color"
	"errors"
)

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	angle = math.Pi / 6
)

var red, blue = color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func svg() string {
	_, _, high, _ := corner(0, 0)
	low := high
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z, _ := corner(i, j)
			high = math.Max(high, z)
			low = math.Min(low, z)
		}
	}
	s := fmt.Sprintf(
		"<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, bz, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, cz,err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, dz, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			z := az * 0.25 + bz * 0.25 + cz * 0.25 + dz * 0.25
			strokeColor := colorToHex(interpolateColor(red, blue, (z-low)/(high-low)))
			s += fmt.Sprintf("<polygon stroke=\"%s\" points='%g,%g %g,%g %g,%g %g,%g' />\n",
				strokeColor, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	s += fmt.Sprintln("</svg>")
	return s
}

func lerp(a uint8, b uint8, t float64) uint8 {
	min, max := a, b
	if a > b {
		min, max = b, a
	}
	return min + uint8(float64(max-min)*t)
}

func interpolateColor(c1, c2 color.Color, t float64) color.Color {
	// TODO: assert valid t
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return color.RGBA{
		lerp(uint8(r1>>8), uint8(r2>>8), t),
		lerp(uint8(g1>>8), uint8(g2>>8), t),
		lerp(uint8(b1>>8), uint8(b2>>8), t),
		lerp(uint8(a1>>8), uint8(a2>>8), t),
	}
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(r>>8), uint8(g>>8), uint8(b>>8))
}

func f(x, y float64) (float64, error) {
	r := math.Hypot(x, y)

	if r == 0 {
		return 0, errors.New("division by zero")
	}

	result := math.Sin(r) / r

	if math.IsInf(result, 0) {
		return result, errors.New("infinite result")
	}

	if math.IsNaN(result) {
		return result, errors.New("NaN result")
	}

	return result, nil
}

// func f(x, y float64) (float64, error) {
// 	result := x*x - y*y
// 	if math.IsInf(result, 0) {
// 		return result, errors.New("infinite result")
// 	}

// 	if math.IsNaN(result) {
// 		return result, errors.New("NaN result")
// 	}
// 	return result, nil
// }

func corner(i, j int) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z,err := f(x, y)
	if err != nil {
		return 0, 0, 0, err
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func main() {
	s := svg()
	fmt.Println(s)
}