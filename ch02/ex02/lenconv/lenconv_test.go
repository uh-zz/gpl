package lenconv

import (
	"testing"
)

func TestMToF(t *testing.T) {
	var expected Feet = 20
	if MToF(6) != expected {
		t.Errorf("got: %v want: %v", MToF(6), expected)
	}
}

func TestFToM(t *testing.T) {
	var expected Meters = 1.5
	if FToM(5) != expected {
		t.Errorf("got: %v want: %v", FToM(5), expected)
	}
}
