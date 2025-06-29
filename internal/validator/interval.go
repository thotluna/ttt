package validator

type Interval struct {
	min int
	max int
}

func NewInterval(min, max int) *Interval {
	return &Interval{
		min: min,
		max: max,
	}
}

func (i *Interval) Contains(value int) bool {
	return value >= i.min && value <= i.max
}

func (i *Interval) Min() int {
	return i.min
}

func (i *Interval) Max() int {
	return i.max
}
