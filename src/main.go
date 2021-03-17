package main

import (
	"bufio"
	"fmt"
	"math"
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
	a := Player{
		n1:                    n[0],
		n2:                    n[1],
		n3:                    n[2],
		minKey:                0,
		maxKey:                0,
		publication:           map[int]rational.Rational{},
		cumulativePublication: map[int]rational.Rational{},
	}.init()
	p := Player{
		n1:                    m[0],
		n2:                    m[1],
		n3:                    m[2],
		minKey:                0,
		maxKey:                0,
		publication:           map[int]rational.Rational{},
		cumulativePublication: map[int]rational.Rational{},
	}.init()
	// a.init()
	// p.init()

	// fmt.Println(a.value)
	// fmt.Println(a.sumValue)
	// fmt.Println(p.value)
	// fmt.Println(p.sumValue)
	fmt.Println("The probability that the an active player will win: ")
	result := a.buttle(p)
	fmt.Printf("%d/%d, %.2f%%\n", result.GetNumerator(), result.GetDenominator(), result.Float64()*100)

}

//  Player プレイヤー構造体。必ずinitメソッドを呼んでください
type Player struct {
	n1                    int
	n2                    int
	n3                    int
	minKey                int
	maxKey                int
	publication           map[int]rational.Rational
	cumulativePublication map[int]rational.Rational
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

func (p Player) init() Player {
	ls := p.roll_sum()
	denominator := len(ls)

	for i := 0; i < denominator; i++ {
		_, ok := p.publication[ls[i]]
		if ok {
			p.publication[ls[i]] = p.publication[ls[i]].Add(rational.New(1, int64(denominator)))
		} else {
			p.publication[ls[i]] = rational.New(1, int64(denominator))
		}
	}

	tmpMin := math.MaxInt32
	tmpMax := 0
	for key := range p.publication {
		if tmpMin > key {
			tmpMin = key
		}
		if tmpMax < key {
			tmpMax = key
		}
	}
	p.minKey = tmpMin
	p.maxKey = tmpMax

	p.cumulativePublication[p.minKey] = p.publication[p.minKey]
	for i := p.minKey + 1; i <= p.maxKey; i++ {
		p.cumulativePublication[i] = p.cumulativePublication[i-1].Add(p.publication[i])
	}

	return p
}

func (active Player) buttle(pussive Player) (result rational.Rational) {
	for i := active.minKey; i <= active.maxKey; i++ {
		var pub rational.Rational
		if i < pussive.minKey {
			continue
		}
		if pussive.maxKey < i {
			pub = rational.New(1, 1)
		} else {
			pub = pussive.cumulativePublication[i]
		}
		result = result.Add(pub.Multiply(active.publication[i]))
	}
	return
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
