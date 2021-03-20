package player

import (
	"math"

	"github.com/alex-ant/gomath/rational"
)

//  Player プレイヤー構造体。必ずinitメソッドを呼んでください
type Player struct {
	isRoll                bool
	N1                    int
	N2                    int
	N3                    int
	MinKey                int
	MaxKey                int
	Publication           map[int]rational.Rational
	CumulativePublication map[int]rational.Rational
}

// func (p Player) roll() (result [][]int) {
// 	var tmp [][]int
// 	for i := 1; i <= p.N2; i++ {
// 		tmp = append(tmp, []int{i})
// 	}
// 	// 追加用
// 	value := tmp
// 	for i := 1; i < p.N1; i++ {
// 		lenResult := len(value)
// 		for j := 0; j < lenResult; j++ {
// 			for k := 0; k < len(tmp); k++ {
// 				value = append(value, append(value[j], tmp[k]...))
// 			}
// 		}
// 		pow := int(math.Pow(float64(p.N2), float64(i)))
// 		value = value[pow:]
// 	}
// 	return value
// }

// func product(n, m int) {
func (p Player) product() (result [][]int) {
	// n d m の試行を考える
	n := p.N1
	m := p.N2
	result = make([][]int, int(math.Pow(float64(m), float64(n))))

	// m^0, m^1, ... m^n となる配列を作成する
	expoList := make([]int, n)
	expoList[0] = 1
	for i := 1; i < n; i++ {
		expoList[i] = expoList[i-1] * m
	}

	for i := 0; i < int(math.Pow(float64(m), float64(n))); i++ {
		result[i] = baseTrance(i, n, m, expoList)
	}
	return result
}

func (p Player) roll_sum() (result []int) {
	// m := p.roll()
	m := p.product()
	result = make([]int, len(m))
	for i := 0; i < len(m); i++ {
		sum := 0
		for j := 0; j < len(m[i]); j++ {
			sum += m[i][j]
		}
		result[i] = sum + p.N3
	}
	return
}

func (p Player) Init() Player {

	if p.N1 == 0 {
		p.isRoll = false
		p.MinKey = p.N3
		p.MaxKey = p.N3
		return p
	}

	p.isRoll = true

	ls := p.roll_sum()
	denominator := len(ls)

	for i := 0; i < denominator; i++ {
		_, ok := p.Publication[ls[i]]
		if ok {
			p.Publication[ls[i]] = p.Publication[ls[i]].Add(rational.New(1, int64(denominator)))
		} else {
			p.Publication[ls[i]] = rational.New(1, int64(denominator))
		}
	}

	tmpMin := math.MaxInt32
	tmpMax := 0
	for key := range p.Publication {
		if tmpMin > key {
			tmpMin = key
		}
		if tmpMax < key {
			tmpMax = key
		}
	}
	p.MinKey = tmpMin
	p.MaxKey = tmpMax

	p.CumulativePublication[p.MinKey] = p.Publication[p.MinKey]
	for i := p.MinKey + 1; i <= p.MaxKey; i++ {
		p.CumulativePublication[i] = p.CumulativePublication[i-1].Add(p.Publication[i])
	}

	return p
}

type pat struct {
	a bool
	p bool
}

func (active Player) Buttle(pussive Player) (result rational.Rational) {
	// 判定の種類として
	// Roll  vs Roll
	// Roll  vs Const
	// Const vs Roll
	// Const vs Const
	// の4種類がある．現在実装されているのは，R vs R のみ．
	// Buttle を これらを呼び出すメソッドとして実装する

	p := pat{
		pussive.isRoll,
		active.isRoll,
	}
	switch p {
	case pat{true, true}:
		result = pussive.rvsr(active)
	case pat{true, false}:
		result = pussive.rvsc(active)
	case pat{false, true}:
		result = active.rvsc(pussive)
	case pat{false, false}:
		result = pussive.cvsc(active)
	}

	return
}

func (active Player) rvsr(pussive Player) (result rational.Rational) {
	for i := active.MinKey; i <= active.MaxKey; i++ {
		var pub rational.Rational
		if i < pussive.MinKey {
			continue
		}
		if pussive.MaxKey < i {
			pub = rational.New(1, 1)
		} else {
			pub = pussive.CumulativePublication[i]
		}
		result = result.Add(pub.Multiply(active.Publication[i]))
	}
	return
}

// rvsc -> Roll vs Const
func (active Player) rvsc(pussive Player) (result rational.Rational) {
	// fmt.Printf("%+v\n", pussive)
	// fmt.Printf("%+v\n", active)
	return rational.New(1, 1).Subtract(active.CumulativePublication[pussive.N3-1])
}

// cvsc -> Const vs Const
func (active Player) cvsc(pussive Player) rational.Rational {
	if pussive.N3 <= active.N3 {
		return rational.New(0, 1)
	}
	return rational.New(1, 1)
}

func baseTrance(x, len, base int, expoList []int) []int {
	result := make([]int, len)
	for i := len - 1; i >= 0; i-- {
		result[i] = x / expoList[i]
		x -= result[i] * expoList[i]
		result[i]++
	}
	// fmt.Println(result)
	return result
}
