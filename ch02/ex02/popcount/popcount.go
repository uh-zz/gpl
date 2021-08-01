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
//
// [pattern 1] テーブルを参照するタイプ
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

// [pattern 2] 1bitずつ見ていくパターン
// これやってたわ。。(一番効率悪)
func PopCountByOneBit(x uint64) int {
	var (
		cnt  int
		mask uint64 = 1
	)
	for ; mask != 0; mask = mask << 1 {
		if x&mask > 0 {
			cnt++
		}
	}

	return cnt
}
