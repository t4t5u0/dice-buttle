package player

import (
	"testing"

	"github.com/alex-ant/gomath/rational"
)

func TestButtle(t *testing.T) {
	// 必要なテストケース
	// r vs r
	// 2d6 / 2d6
	// 2d6 + 1 / 2d6
	// 2d6 / 2d6 + 1
	// 2d6 + 3 / 2d6 + 1
	// 2d6 / 2d6 + 13
	// 2d6 - 1 / 2d6
	//
	// 2d6 / 2d10
	//
	// r vs c
	// 2d6 / 1
	// 2d6 / 8
	// 2d6 / 12
	// 2d6 / 13

	// c vs r
	// 1 / 2d6
	// 8 / 2d6
	// 11 / 2d6
	// 12 / 2d6
	// 13 / 2d6

	// c vs c
	// 3 / 3
	// 1 / 3
	// 3 / 1

	// 10d6 // 激遅になる．6^10 = 60_466_176 回の計算が必要になるから仕方ない

	activeAList := []int{2}
	activeXList := []int{}
	activeNList := []int{}

	pussiveAList := []int{2}
	pussiveXList := []int{}
	pussiveNList := []int{}

	activeList := make([]Player, len(activeAList))
	pussiveList := make([]Player, len(activeAList))
	for i := 0; i < len(activeAList); i++ {
		activeList[i] = Player{
			isRoll:                false,
			N1:                    activeAList[i],
			N2:                    activeXList[i],
			N3:                    activeNList[i],
			MinKey:                0,
			MaxKey:                0,
			Publication:           map[int]rational.Rational{},
			CumulativePublication: map[int]rational.Rational{},
		}.Init()
		pussiveList[i] = Player{
			isRoll:                false,
			N1:                    pussiveAList[i],
			N2:                    pussiveXList[i],
			N3:                    pussiveNList[i],
			MinKey:                0,
			MaxKey:                0,
			Publication:           map[int]rational.Rational{},
			CumulativePublication: map[int]rational.Rational{},
		}.Init()
	}
}
