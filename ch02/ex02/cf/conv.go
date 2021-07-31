package main

import (
	"fmt"

	"github.com/uh-zz/gpl/ch02/ex01/tempconv"
	"github.com/uh-zz/gpl/ch02/ex02/lenconv"
)

func ArgsToTempConv(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func ArgsToLenConv(t float64) {
	m := lenconv.Meters(t)
	ft := lenconv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToF(m), ft, lenconv.FToM(ft))
}
