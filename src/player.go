package player

import (
	"fmt"
	"math"
	"strconv"

	"github.com/alex-ant/gomath/rational"
)

//  Player プレイヤー構造体。必ずinitメソッドを呼んでください
type Player struct {
	N1                    int
	N2                    int
	N3                    int
	MinKey                int
	MaxKey                int
	Publication           map[int]rational.Rational
	CumulativePublication map[int]rational.Rational
}

func (p Player) roll() (result [][]int) {
	var tmp [][]int
	for i := 1; i <= p.N2; i++ {
		tmp = append(tmp, []int{i})
	}
	// 追加用
	value := tmp
	for i := 1; i < p.N1; i++ {
		lenResult := len(value)
		for j := 0; j < lenResult; j++ {
			for k := 0; k < len(tmp); k++ {
				value = append(value, append(value[j], tmp[k]...))
			}
		}
		pow := int(math.Pow(float64(p.N2), float64(i)))
		value = value[pow:]
	}
	return value
}

// func product(n, m int) {
func product() {
	// n d m の試行を考える
	n := 10
	m := 6
	for i := 0; i < int(math.Pow(float64(m), float64(n))); i++ {
		a := []rune(fmt.Sprintf("%0"+fmt.Sprintf("%d", n)+"s", strconv.FormatInt(int64(i), m)))
		b := make([]int, n)
		for j, item := range a {
			b[j], _ = strconv.Atoi(string(item))
			b[j]++
		}
		// fmt.Println(b)
	}
}

func (p Player) roll_sum() (result []int) {
	m := p.roll()
	for i := 0; i < len(m); i++ {
		sum := 0
		for j := 0; j < len(m[i]); j++ {
			sum += m[i][j]
		}
		result = append(result, sum+p.N3)
	}
	return
}

func (p Player) Init() Player {
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

func (active Player) Buttle(pussive Player) (result rational.Rational) {
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
