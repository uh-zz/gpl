package main

import "fmt"

func main() {

	data := []string{"one", "one", "two", "three", "three", "three"}

	fmt.Println(deduplication(data))
}

func deduplication(slice []string) []string {

	for i, _ := range slice {
		slice = compare(i, slice, "left compare")
		slice = compare(i+1, slice, "right compare")
	}
	return slice
}

func compare(index int, slice []string, option string) []string {
	fmt.Println(len(slice), index, slice)
	if len(slice) <= index+1 {
		return slice
	}

	left := slice[index]
	right := slice[index+1]

	if left == right {
		fmt.Println(option, "remove", slice[index+1], "from", slice)
		slice = remove(slice, index+1)
		fmt.Println(option, "removed", slice)
	}

	return slice
}

func remove(slice []string, index int) []string {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
