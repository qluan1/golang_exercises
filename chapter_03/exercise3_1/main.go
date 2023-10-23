package main

import (
	"fmt"
	"math"
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

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	s := fmt.Sprintf(
		"<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j< cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			s += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	s += fmt.Sprintln("</svg>")
	fmt.Println(s)
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z,err := f(x, y)
	if err != nil {
		return 0, 0, err
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
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
