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
	fmt.Println("A d X + N: N is optional")
	fmt.Println("if A is 0, N is const")
	fmt.Print("active  player A X N: ")
	n, err := SplitIntStdin(" ")
	if err != nil || len(n) < 2 {
		println("不正な入力です")
	}
	if len(n) == 2 {
		n = append(n, 0)
	}

	fmt.Print("passive player A X N: ")
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
	fmt.Println("The probability that the an pussive player will win: ")
	result := p.Buttle(a)
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
