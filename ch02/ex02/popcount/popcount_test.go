package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	// 0000 1000 -> 8
	var expected int = 1
	if PopCount(8) != expected {
		t.Errorf("got: %v want: %v", PopCount(8), expected)
	}

	// 0000 0101 -> 5
	expected = 2
	if PopCount(5) != expected {
		t.Errorf("got: %v want: %v", PopCount(5), expected)
	}

	// 0000 1110 -> 14
	expected = 3
	if PopCount(14) != expected {
		t.Errorf("got: %v want: %v", PopCount(14), expected)
	}

	// 0000 1111 -> 15
	expected = 4
	if PopCount(15) != expected {
		t.Errorf("got: %v want: %v", PopCount(15), expected)
	}
}

func TestPopCountByOneBit(t *testing.T) {
	// 0000 1111 -> 15
	var expected int = 4
	if PopCountByOneBit(15) != expected {
		t.Errorf("got: %v want: %v", PopCountByOneBit(15), expected)
	}
}

func TestPopCountOptimize(t *testing.T) {
	// 0000 1111 -> 15
	var expected int = 4
	if PopCountOptimize(15) != expected {
		t.Errorf("got: %v want: %v", PopCountOptimize(15), expected)
	}
}
