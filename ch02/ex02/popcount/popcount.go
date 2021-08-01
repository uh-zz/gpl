package popcount

// pc[i] はiのポピュレーションカウントです
// 解説 -> https://play.golang.org/p/j8LVzz4Hjk9
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount はxのポピュレーションカウント(1が設定されているビット数)を返す
// 著者が言うには、このアルゴリズムはビットを数える最速のアルゴリズムではないらしい
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
