// テストあとから書く
package tempconv

import (
	"fmt"

	"testing"
)

const (
	UnitCelsius    string  = "°C"
	UnitFahrenheit string  = "°F"
	TempValue      float64 = 100
)

var (
	c Celsius    = Celsius(TempValue)
	f Fahrenheit = Fahrenheit(TempValue)
)

func TestToString(t *testing.T) {
	var (
		expected string
	)

	expected = fmt.Sprintf("%v", TempValue) + UnitCelsius
	if c.String() != expected {
		t.Errorf("got: %v want: %v", c.String(), expected)
	}

	expected = fmt.Sprintf("%v", TempValue) + UnitFahrenheit
	if f.String() != expected {
		t.Errorf("got: %v want: %v", f.String(), expected)
	}
}

func TestCToF(t *testing.T) {
	var expected Fahrenheit = 212
	if CToF(c) != expected {
		t.Errorf("got: %v want: %v", CToF(c), expected)
	}
}
func TestFToC(t *testing.T) {
	var expected Celsius = 37.77777777777778
	if FToC(f) != expected {
		t.Errorf("got: %v want: %v", FToC(f), expected)
	}
}

func TestKToC(t *testing.T) {
	var expected Celsius = AbsoluteZeroC
	if KToC(0) != expected {
		t.Errorf("got: %v want: %v", KToC(0), expected)
	}
}

func TestCToK(t *testing.T) {
	var expected Kelvin = Kelvin(AbsoluteZeroC)
	if CToK(0) != -expected {
		t.Errorf("got: %v want: %v", CToK(0), expected)
	}
}

func TestKToF(t *testing.T) {
	var expected Fahrenheit = 459.66999999999996
	if KToF(0) != -expected {
		t.Errorf("got: %v want: %v", KToF(0), expected)
	}
}
