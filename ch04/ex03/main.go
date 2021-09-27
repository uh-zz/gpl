package main

import "fmt"

const ArrayLength = 5

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	b := [ArrayLength]int{0, 1, 2, 3, 4}
	pointerReverse(&b)
	fmt.Println(b)
}

// reverse はintのスライスを直接逆順に並べ替える
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// pointerReverseは配列のポインタを受け取り、逆順に並べ替える
func pointerReverse(s *[ArrayLength]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
