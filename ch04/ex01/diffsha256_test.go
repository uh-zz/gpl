package main

import (
	"crypto/sha256"
	"testing"
)

func TestDiffSha256(t *testing.T) {
	var (
		zeroBytesList [sha256.Size]int
		diffBytesList [sha256.Size]int = [sha256.Size]int{4, 3, 6, 5, 4, 4, 4, 4, 5, 0, 6, 3, 2, 2, 4, 4, 4, 5, 5, 3, 6, 3, 5, 5, 7, 2, 5, 3, 3, 3, 3, 3}
	)

	var tests = []struct {
		pattern string
		c1      [sha256.Size]byte
		c2      [sha256.Size]byte
		out     [sha256.Size]int
	}{
		{
			"same",
			sha256.Sum256([]byte("x")),
			sha256.Sum256([]byte("x")),
			zeroBytesList,
		},
		{
			"diff",
			sha256.Sum256([]byte("x")),
			sha256.Sum256([]byte("X")),
			diffBytesList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			got := DiffSha256(tt.c1, tt.c2)

			if got != tt.out {
				t.Errorf("got: %v\n\t\t\twant: %v", got, tt.out)
			}
		})
	}
}
