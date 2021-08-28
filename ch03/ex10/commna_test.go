package main

import (
	"testing"
)

// 確保した容量とかアロケーション数はほぼ同じだけど実行時間が2倍も変わっちゃう。。
//
//
// go test -bench . -benchmem
// BenchmarkComma-4                24444198                48.14 ns/op            8 B/op          1 allocs/op
// BenchmarkBufComma-4              5541354               183.3 ns/op            80 B/op          2 allocs/op
// BenchmarkWithDefaultBuffer-4    10894846                97.68 ns/op           72 B/op          2 allocs/op
//
// reverseしてる分でした泣
//
// reverseしなかった場合
// go test -bench . -benchmem
// BenchmarkComma-4                23949460                50.20 ns/op            8 B/op          1 allocs/op
// BenchmarkBufComma-4             11657074                95.08 ns/op           72 B/op          2 allocs/op
// BenchmarkWithDefaultBuffer-4    11871906                92.69 ns/op           72 B/op          2 allocs/op

const numberStr = "12345"

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {

		comma(numberStr)
	}
}

func BenchmarkBufComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bufcomma(numberStr)
	}
}

func BenchmarkWithDefaultBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithDefaultBuffer(numberStr)
	}
}
