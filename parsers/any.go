package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func NewAny(val string, minAndMaxGetter timeunits.MinAndMaxGetter) Any {
	return Any{
		val,
		minAndMaxGetter,
	}
}

type Any struct {
	val             string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s Any) Parse() (string, error) {
	min := s.minAndMaxGetter.Min()
	max := s.minAndMaxGetter.Max()
	var output string
	for i := min; i <= max; i++ {
		if i == min {
			output = fmt.Sprintf("%v", i)
		} else {
			output = fmt.Sprintf("%v %v", output, i)
		}
	}

	return output, nil
}
