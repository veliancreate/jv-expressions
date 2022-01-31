package timeunits

type MinAndMaxGetter interface {
	Min() int
	Max() int
}

type Unit struct {
	min int
	max int
}

func (u Unit) Min() int {
	return u.min
}

func (u Unit) Max() int {
	return u.max
}

func NewMinute() Unit {
	return Unit{
		min: 0,
		max: 59,
	}
}

func NewHour() Unit {
	return Unit{
		min: 0,
		max: 23,
	}
}

func NewMonth() Unit {
	return Unit{
		min: 1,
		max: 12,
	}
}

func NewDayOfWeek() Unit {
	return Unit{
		min: 1,
		max: 7,
	}
}

func NewDayOfMonth() Unit {
	return Unit{
		min: 1,
		max: 31,
	}
}
