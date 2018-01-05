package roll

import (
	"math"
	"math/rand"
	"sort"
)

type Dice struct {
	Sides int
	Num   int
	Rolls []int
}

func NewDice(sides int, num int) *Dice {
	rolls := make([]int, int(math.Abs(float64(num))))
	d := Dice{
		Sides: sides,
		Num:   num,
		Rolls: rolls,
	}
	d.roll()
	return &d
}

func (d *Dice) roll() {
	for i := 0; i < int(math.Abs(float64(d.Num))); i++ {
		d.Rolls[i] = rand.Intn(d.Sides) + 1
	}
}

func (d *Dice) Total() int {
	if d.Sides == 20 && d.Num == 2 {
		return d.max()
	} else if d.Sides == 20 && d.Num == -2 {
		return d.min()
	}
	return d.total()
}

func (d *Dice) total() int {
	var t int
	for _, r := range d.Rolls {
		t += r
	}
	return t
}

func (d *Dice) max() int {
	sort.Ints(d.Rolls)
	return d.Rolls[1]
}

func (d *Dice) min() int {
	sort.Ints(d.Rolls)
	return d.Rolls[0]
}
