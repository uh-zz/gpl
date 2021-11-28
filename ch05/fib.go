package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	fmt.Println(n, n-2, n-1)
	return fib(n-2) + fib(n-1)
}
