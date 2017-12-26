package roll

// Dice is an interface for all the
// possible dice type
type Dice interface {
	roll()
	total() int
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

func (d *D6) total() int {
	var s int
	for r := range d.rolls {
		s += r
	}
	return s
}
