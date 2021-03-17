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
	// r := roll(n[0], n[1])
	// fmt.Printf("%dd%d = %v\n", n[0], n[1], r)
	// fmt.Println(roll_sum(n[0], n[1], n[2]))
	// m := multiset{
	// 	m: map[string]int{},
	// }
	// m.set(roll_sum(n[0], n[1], n[2]))
	// fmt.Println(m)
	a := Player{isActive: false,
		n1: n[0],
		n2: n[1],
		n3: n[2],
		sumValue: multiset{
			map[int]int{},
		}}
	a.set()
	fmt.Println(a.sumValue)

}

type Player struct {
	isActive bool
	n1       int
	n2       int
	n3       int
	sumValue multiset
}

func (p Player) roll() (result [][]int) {
	var tmp [][]int
	for i := 1; i <= p.n2; i++ {
		tmp = append(tmp, []int{i})
	}
	// 追加用
	value := tmp
	for i := 1; i < p.n1; i++ {
		lenResult := len(value)
		for j := 0; j < lenResult; j++ {
			for k := 0; k < len(tmp); k++ {
				value = append(value, append(value[j], tmp[k]...))
			}
		}
		pow := int(math.Pow(float64(p.n2), float64(i)))
		value = value[pow:]
	}
	return value
}

func (p Player) roll_sum() (result []int) {
	m := p.roll()
	for i := 0; i < len(m); i++ {
		sum := 0
		for j := 0; j < len(m[i]); j++ {
			sum += m[i][j]
		}
		result = append(result, sum+p.n3)
	}
	return
}

type multiset struct {
	ms map[int]int
}

func (p Player) set() {
	ls := p.roll_sum()
	for i := 0; i < len(ls); i++ {
		_, ok := p.sumValue.ms[ls[i]]
		if ok {
			p.sumValue.ms[ls[i]] += 1
		} else {
			p.sumValue.ms[ls[i]] = 1
		}
	}
}

func (p Player) buttle() {

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
