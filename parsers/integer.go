package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newInteger(val int, minAndMaxGetter timeunits.MinAndMaxGetter) (integer, error) {
	max := minAndMaxGetter.Max()
	min := minAndMaxGetter.Min()

	if val > max {
		return integer{}, fmt.Errorf("val %v is over max %v", val, max)
	}

	if val < min {
		return integer{}, fmt.Errorf("val %v is under min %v", val, min)
	}

	return integer{
		value: fmt.Sprintf("%v", val),
	}, nil
}

type integer struct {
	value string
}

func (i integer) Value() string {
	return i.value
}
