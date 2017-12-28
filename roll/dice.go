package roll

import (
	"math/rand"
)

type Dice struct {
	Sides int
	Num   int
	Rolls []int
}

func NewDice(sides int, num int) *Dice {
	rolls := make([]int, num)
	d := Dice{
		Sides: sides,
		Num:   num,
		Rolls: rolls,
	}
	d.roll()
	return &d
}

func (d *Dice) roll() {
	for i := 0; i < d.Num; i++ {
		d.Rolls[i] = rand.Intn(d.Sides) + 1
	}
}

func (d *Dice) Total() int {
	var t int
	for _, r := range d.Rolls {
		t += r
	}
	return t
}
