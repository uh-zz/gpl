// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2015 Yoshiki Shibata
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			value := 255 - contrast*n
			return color.RGBA{value, contrast * n, 0, 0xff}
		}
	}
	return color.RGBA{0, 0, 0, 0xff} // black
}

func main() {
	const (
		width  = 1024
		height = 1024
		scale  = 4
	)

	handler := func(w http.ResponseWriter, r *http.Request) {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := scale * (float64(py)/height - 0.5)
			for px := 0; px < width; px++ {
				x := scale * (float64(px)/width - 0.5)
				z := complex(x, y)
				// Image point (px, py) represents complex value z
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(w, img)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
