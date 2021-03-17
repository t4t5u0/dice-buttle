package main

import (
	"bufio"
	player "dice-buttle/src"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alex-ant/gomath/rational"
)

func main() {
	// m1 := multiset.Multiset{"a": 1, "b": 2}
	// fmt.Println(multiset.Sum(m1))
	// n1 d n2 + n3
	// m1 d m2 + m3
	fmt.Println("n1 d n2 + n3: n3 is optional")
	fmt.Print("active  player n1 n2 n3: ")
	n, err := SplitIntStdin(" ")
	if err != nil || len(n) < 2 {
		println("不正な入力です")
	}
	if len(n) == 2 {
		n = append(n, 0)
	}

	fmt.Print("passive player m1 m2 m3: ")
	m, err := SplitIntStdin(" ")
	if err != nil || len(m) < 2 {
		println("不正な入力です")
	}
	if len(m) == 2 {
		m = append(m, 0)
	}
	a := player.Player{
		N1:                    n[0],
		N2:                    n[1],
		N3:                    n[2],
		MinKey:                0,
		MaxKey:                0,
		Publication:           map[int]rational.Rational{},
		CumulativePublication: map[int]rational.Rational{},
	}.Init()
	p := player.Player{
		N1:                    m[0],
		N2:                    m[1],
		N3:                    m[2],
		MinKey:                0,
		MaxKey:                0,
		Publication:           map[int]rational.Rational{},
		CumulativePublication: map[int]rational.Rational{},
	}.Init()
	// a.init()
	// p.init()

	// fmt.Println(a.value)
	// fmt.Println(a.sumValue)
	// fmt.Println(p.value)
	// fmt.Println(p.sumValue)
	fmt.Println("The probability that the an active player will win: ")
	result := a.Buttle(p)
	fmt.Printf("%d/%d, %.2f%%\n", result.GetNumerator(), result.GetDenominator(), result.Float64()*100)

}

// SplitWithoutEmpty 入力から空白をトリムするやつ
func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}

	return
}

// 文字列を1行入力
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

// デリミタで分割して文字列スライスを取得
func SplitStrStdin(delim string) (stringReturned []string) {
	// 文字列で1行取得
	stringInput := StrStdin()

	// 空白で分割
	// stringSplited := strings.Split(stringInput, delim)
	stringReturned = SplitWithoutEmpty(stringInput, delim)

	return
}

func SplitIntStdin(delim string) (intSlices []int, err error) {
	// 文字列スライスを取得
	stringSplited := SplitStrStdin(" ")

	// 整数スライスに保存
	for i := range stringSplited {
		var iparam int
		iparam, err = strconv.Atoi(stringSplited[i])
		if err != nil {
			return
		}
		intSlices = append(intSlices, iparam)
	}
	return
}
