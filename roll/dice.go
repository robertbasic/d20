package roll

import (
	"math/rand"
)

// Dice is an interface for all the
// possible dice type
type Dice interface {
	roll()
	Total() int
}

// D20 is a Dice with 20 sides
type D20 struct {
	sides        int
	num          int
	advantage    bool
	disadvantage bool
	rolls        []int
}

func NewD20Normal() *D20 {
	rolls := make([]int, 1)
	d := &D20{
		num:          1,
		advantage:    false,
		disadvantage: false,
		sides:        20,
		rolls:        rolls,
	}
	d.roll()
	return d
}

func NewD20Advantage() *D20 {
	rolls := make([]int, 2)
	d := &D20{
		num:          2,
		advantage:    true,
		disadvantage: false,
		sides:        20,
		rolls:        rolls,
	}
	d.roll()
	return d
}

func NewD20Disadvantage() *D20 {
	rolls := make([]int, 2)
	d := &D20{
		num:          2,
		advantage:    false,
		disadvantage: true,
		sides:        20,
		rolls:        rolls,
	}
	d.roll()
	return d
}

func (d *D20) roll() {

}

func (d *D20) Total() int {
	return d.rolls[0]
}

// D6 is a Dice with 6 sides
type D6 struct {
	sides int
	num   int
	rolls []int
}

// NewD6 rolls n number of new 6 sided Dice
func NewD6(n int) *D6 {
	rolls := make([]int, n)
	d := &D6{
		num:   n,
		sides: 6,
		rolls: rolls,
	}
	d.roll()
	return d
}

func (d *D6) roll() {
	for i := 0; i < d.num; i++ {
		d.rolls[i] = rand.Intn(d.sides) + 1
	}
}

func (d *D6) Total() int {
	var s int
	for _, r := range d.rolls {
		s += r
	}
	return s
}

// D4 is a Dice with 4 sides
type D4 struct {
	sides int
	num   int
	rolls []int
}

// NewD4 rolls n number of new 4 sided Dice
func NewD4(n int) *D4 {
	rolls := make([]int, n)
	d := &D4{
		num:   n,
		sides: 4,
		rolls: rolls,
	}
	d.roll()
	return d
}

func (d *D4) roll() {

}

func (d *D4) Total() int {
	var s int
	for _, r := range d.rolls {
		s += r
	}
	return s
}

// D8 is a Dice with 8 sides
type D8 struct {
	sides int
	num   int
	rolls []int
}

// NewD8 rolls n number of new 8 sided Dice
func NewD8(n int) *D8 {
	rolls := make([]int, n)
	d := &D8{
		num:   n,
		sides: 8,
		rolls: rolls,
	}
	d.roll()
	return d
}

func (d *D8) roll() {

}

func (d *D8) Total() int {
	var s int
	for _, r := range d.rolls {
		s += r
	}
	return s
}

// D10 is a Dice with 10 sides
type D10 struct {
	sides int
	num   int
	rolls []int
}

// NewD10 rolls n number of new 10 sided Dice
func NewD10(n int) *D10 {
	rolls := make([]int, n)
	d := &D10{
		num:   n,
		sides: 10,
		rolls: rolls,
	}
	d.roll()
	return d
}

func (d *D10) roll() {

}

func (d *D10) Total() int {
	var s int
	for _, r := range d.rolls {
		s += r
	}
	return s
}

// D12 is a Dice with 12 sides
type D12 struct {
	sides int
	num   int
	rolls []int
}

// NewD12 rolls n number of new 12 sided Dice
func NewD12(n int) *D12 {
	rolls := make([]int, n)
	d := &D12{
		num:   n,
		sides: 12,
		rolls: rolls,
	}
	d.roll()
	return d
}

func (d *D12) roll() {

}

func (d *D12) Total() int {
	var s int
	for _, r := range d.rolls {
		s += r
	}
	return s
}

// D100 is a Dice with 100 sides
type D100 struct {
	sides int
	num   int
	rolls []int
}

// NewD100 rolls n number of new 100 sided Dice
func NewD100(n int) *D100 {
	rolls := make([]int, n)
	d := &D100{
		num:   n,
		sides: 100,
		rolls: rolls,
	}
	d.roll()
	return d
}

func (d *D100) roll() {

}

func (d *D100) Total() int {
	var s int
	for _, r := range d.rolls {
		s += r
	}
	return s
}
