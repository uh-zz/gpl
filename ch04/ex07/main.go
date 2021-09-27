package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abcdef"
	b := []byte(s)
	ReverseUTF8(b)
	fmt.Println(string(b))
}

func reverseByte(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func rotateBytes(b []byte, p int) {
	reverseByte(b[:p])
	reverseByte(b[p:])
	reverseByte(b)
}

func ReverseUTF8(b []byte) {

	if len(b) == 0 {
		return
	}

	r, size := utf8.DecodeRune(b)

	if r == utf8.RuneError {
		panic("Rune Error")
	}

	rotateBytes(b, size)
	fmt.Println("after", string(b))

	// 再帰的に、bのルーン１文字を引いたバイトスライスを渡す
	// こうすることで最後の１文字までrotateBytesを繰り返す
	// 0文字になるとreturnしてくれる
	// ReverseUTF8を返すたびに、fmt.Println("res", string(b), len(b))を実行する
	// →最後の出力以外は途中計算なので、無視しておｋ
	ReverseUTF8(b[:len(b)-size])

	fmt.Println("res", string(b), len(b))
}
