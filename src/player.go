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

func (active Player) Buttle(passive Player) (result rational.Rational) {
	// 判定の種類として
	// Roll  vs Roll
	// Roll  vs Const
	// Const vs Roll
	// Const vs Const
	// の4種類がある．現在実装されているのは，R vs R のみ．
	// Buttle を これらを呼び出すメソッドとして実装する

	p := pat{
		passive.isRoll,
		active.isRoll,
	}
	switch p {
	case pat{true, true}:
		result = passive.rvsr(active)
	case pat{true, false}:
		result = passive.rvsc(active)
	case pat{false, true}:
		result = active.rvsc(passive)
	case pat{false, false}:
		result = passive.cvsc(active)
	}

	return
}

func (passive Player) rvsr(active Player) (result rational.Rational) {
	// fmt.Println(passive.CumulativePublication)
	// fmt.Println(active.Publication)
	// fmt.Printf("%#v\n", passive)
	// fmt.Printf("%#v\n", active)
	result = rational.New(0, 1)
	for i := passive.MinKey; i <= passive.MaxKey; i++ {
		var pub rational.Rational
		if i < active.MinKey {
			continue
		}
		if active.MaxKey < i {
			pub = rational.New(1, 1)
		} else {
			pub = active.CumulativePublication[i]
		}
		result = result.Add(pub.Multiply(passive.Publication[i]))
	}
	return result
}

// rvsc -> Roll vs Const
func (active Player) rvsc(passive Player) (result rational.Rational) {
	return rational.New(1, 1).Subtract(active.CumulativePublication[passive.N3-1])
}

// cvsc -> Const vs Const
func (active Player) cvsc(passive Player) rational.Rational {
	if passive.N3 <= active.N3 {
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
	return result
}
