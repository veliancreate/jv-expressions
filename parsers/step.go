package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func NewStep(val string, minAndMaxGetter timeunits.MinAndMaxGetter) Step {
	return Step{
		val,
		minAndMaxGetter,
	}
}

type Step struct {
	val             string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s Step) Parse() (string, error) {
	var output string
	min := s.minAndMaxGetter.Min()
	max := s.minAndMaxGetter.Max()

	value, err := stringToInt(s.val, s.minAndMaxGetter)
	if err != nil {
		return "", fmt.Errorf("could not convert step value to int due to %w", err)
	}

	for i := min; i <= max; i += value {
		if i == min {
			output = fmt.Sprintf("%v", i)
		} else {
			output = fmt.Sprintf("%v %v", output, i)
		}
	}

	return output, nil
}
