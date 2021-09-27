package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 2))

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("%v\n", data)
	RotateCycles(data, 7)
	fmt.Printf("%v\n", data)
}

func rotate(slice []int, index int) []int {
	var (
		res    []int
		first  []int = slice[:index]
		second []int = slice[index:]
	)

	res = append(res, second...)
	res = append(res, first...)
	return res
}
