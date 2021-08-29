package main

import (
	"crypto/sha256"
)

func DiffSha256(c1, c2 [32]uint8) [32]int {

	var (
		// バイトごとの異なるビットの数
		diffBits [sha256.Size]int
	)

	for i := range diffBits {
		diffBits[i] = PopCountOptimize(c1[i] ^ c2[i])
	}

	return diffBits
}

func PopCountOptimize(x byte) int {
	var cnt int
	for ; x != 0; x &= x - 1 {
		cnt++
	}
	return cnt
}
