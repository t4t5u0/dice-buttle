package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

	// fmt.Print("passive player m1 m2 m3: ")
	// m, err := SplitIntStdin(" ")
	// if err != nil || len(m) < 2 {
	// 	println("不正な入力です")
	// }
	// if len(m) == 2 {
	// 	m = append(m, 0)
	// }
	// fmt.Println(n, m)
	r := roll(n[0], n[1])
	fmt.Printf("%dd%d = %v\n", n[0], n[1], r)
	fmt.Println(roll_sum(n[0], n[1], n[2]))
	m := multiset{
		m: map[string]int{},
	}
	m.set(roll_sum(n[0], n[1], n[2]))
	fmt.Println(m)

}

func roll(n1, n2 int) (result [][]int) {
	var tmp [][]int
	for i := 1; i <= n2; i++ {
		tmp = append(tmp, []int{i})
	}
	// 追加用
	value := tmp
	for i := 1; i < n1; i++ {
		lenResult := len(value)
		for j := 0; j < lenResult; j++ {
			for k := 0; k < len(tmp); k++ {
				value = append(value, append(value[j], tmp[k]...))
			}
		}
		pow := int(math.Pow(float64(n2), float64(i)))
		value = value[pow:]
	}
	return value
}

// n回直積を取って、足し算して、multiset に突っ込む。母数はn2 ** n1
// 確率にしたさはある
func roll_sum(n1, n2, n3 int) (result []int) {
	m := roll(n1, n2)
	for i := 0; i < len(m); i++ {
		sum := 0
		for j := 0; j < len(m[i]); j++ {
			sum += m[i][j]
		}
		result = append(result, sum+n3)
	}
	return
}

type multiset struct {
	m map[string]int
}

func (d multiset) set(ls []int) {
	for i := 0; i < len(ls); i++ {
		_, ok := d.m[strconv.Itoa(ls[i])]
		if ok {
			d.m[strconv.Itoa(ls[i])] += 1
		} else {
			d.m[strconv.Itoa(ls[i])] = 1
		}
	}
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
