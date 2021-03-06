// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015, 2020 Yoshiki Shibata
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// areAllFinite は各要素がすべて有限なfloat64の場合 true
			if areAllFinite(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func areAllFinite(values ...float64) bool {
	// isFinite
	// true 有限なfloat64
	// false 有限ではないfloat64(無限大もしくはNaN)
	isFinite := func(f float64) bool {
		// math.IsInf(f, sign)
		// sign == 0　の場合、fが無限大かどうかのチェック
		return !math.IsInf(f, 0) && !math.IsNaN(f)
	}

	for _, v := range values {
		// 有限でないfloat64の場合は false
		if !isFinite(v) {
			return false
		}
	}
	return true
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
