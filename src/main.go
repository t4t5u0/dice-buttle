package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/soniakeys/multiset"
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
	fmt.Println(n, m)

}

func roll(n2 int) (result []int) {
	for i := 1; i <= n2; i++ {
		result = append(result, i)
	}
	return
}

func hoge(n1, n2 int) (result multiset.Multiset) {
	r := roll(n2)
	result := multiset.Multiset()
}

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
