package roll

// Dice is an interface for all the
// possible dice type
type Dice interface {
	roll()
	Total() int
}

// D20 is a Dice with 20 sides
type D20 struct {
	advantage    bool
	disadvantage bool
	rolls        []int
}

func NewD20Normal() *D20 {
	d := &D20{advantage: false, disadvantage: false}
	d.roll()
	return d
}

func NewD20Advantage() *D20 {
	d := &D20{advantage: true, disadvantage: false}
	d.roll()
	return d
}

func NewD20Disadvantage() *D20 {
	d := &D20{advantage: false, disadvantage: true}
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
	num   int
	rolls []int
}

// NewD6 rolls n number of new 6 sided Dice
func NewD6(n int) *D6 {
	d := &D6{num: n}
	d.roll()
	return d
}

func (d *D6) roll() {

}

func (d *D6) Total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}

// D4 is a Dice with 4 sides
type D4 struct {
	num   int
	rolls []int
}

// NewD4 rolls n number of new 4 sided Dice
func NewD4(n int) *D4 {
	d := &D4{num: n}
	d.roll()
	return d
}

func (d *D4) roll() {

}

func (d *D4) Total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}

// D8 is a Dice with 8 sides
type D8 struct {
	num   int
	rolls []int
}

// NewD8 rolls n number of new 8 sided Dice
func NewD8(n int) *D8 {
	d := &D8{num: n}
	d.roll()
	return d
}

func (d *D8) roll() {

}

func (d *D8) Total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}

// D10 is a Dice with 10 sides
type D10 struct {
	num   int
	rolls []int
}

// NewD10 rolls n number of new 10 sided Dice
func NewD10(n int) *D10 {
	d := &D10{num: n}
	d.roll()
	return d
}

func (d *D10) roll() {

}

func (d *D10) Total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}

// D12 is a Dice with 12 sides
type D12 struct {
	num   int
	rolls []int
}

// NewD12 rolls n number of new 12 sided Dice
func NewD12(n int) *D12 {
	d := &D12{num: n}
	d.roll()
	return d
}

func (d *D12) roll() {

}

func (d *D12) Total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}

// D100 is a Dice with 100 sides
type D100 struct {
	num   int
	rolls []int
}

// NewD100 rolls n number of new 100 sided Dice
func NewD100(n int) *D100 {
	d := &D100{num: n}
	d.roll()
	return d
}

func (d *D100) roll() {

}

func (d *D100) Total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}
