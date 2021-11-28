package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	// 5 3 4
	// 3 1 2
	// 2 0 1
	// 4 2 3
	// 2 0 1
	// 3 1 2
	// 2 0 1
	// 5
	fmt.Println(fib(5))
}
