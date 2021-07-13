package main

import "fmt"

var global *int

func main() {
	f()
	fmt.Println(*global)
	g()
}

// f uses heap area
// 自由度が高いが少しコストがかかるデータ領域
func f() {
	// [FYI] https://keens.github.io/blog/2017/04/30/memoritosutakkutohi_puto/
	// xはヒープ領域から割り当てられる
	// ↓
	// なぜならローカル変数として宣言されているにもかかわらず、fが戻った後でもglobal変数から到達可能であるから
	// ↓
	// これをxはfからエスケープしているという
	var x int
	x = 1
	global = &x
}

// g uses stack area
// 条件が限られるけど高速に扱えるデータ領域
func g() {
	// *yはgが戻った後では、到達不可能である
	// ↓
	// コンパイラが*yをスタックに割り当てるのは安全
	y := new(int)
	*y = 1
}
