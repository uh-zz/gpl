package main

import "fmt"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
)

func main() {
	fmt.Println(KB, MB, GB, TB)
}
