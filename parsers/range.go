package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func NewRange(val []string, minAndMaxGetter timeunits.MinAndMaxGetter) Range {
	return Range{
		val,
		minAndMaxGetter,
	}
}

type Range struct {
	val             []string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (r Range) Parse() (string, error) {
	var output string

	first, err := stringToInt(r.val[0], r.minAndMaxGetter)
	if err != nil {
		return "", fmt.Errorf("could not convert range value to int due to %w", err)
	}

	last, err := stringToInt(r.val[1], r.minAndMaxGetter)
	if err != nil {
		return "", fmt.Errorf("could not convert range value to int due to %w", err)
	}

	if first > last {
		return "", fmt.Errorf("first range value is higher than last")
	}

	for i := first; i <= last; i++ {
		if i == first {
			output = fmt.Sprintf("%v", i)
		} else {
			output = fmt.Sprintf("%v %v", output, i)
		}
	}

	return output, nil
}
