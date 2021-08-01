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
	var (
		cnt      int
		byteUnit int = 8
	)
	for i := 0; i < byteUnit; i++ {
		cnt += int(pc[byte(x>>(i*8))])
	}
	return cnt
}
