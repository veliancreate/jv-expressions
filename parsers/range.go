package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newRange(val []string, minAndMaxGetter timeunits.MinAndMaxGetter) (ranger, error) {
	var output string

	first, err := stringToInt(val[0], minAndMaxGetter)
	if err != nil {
		return ranger{}, fmt.Errorf("could not convert range value to int due to %w", err)
	}

	last, err := stringToInt(val[1], minAndMaxGetter)
	if err != nil {
		return ranger{}, fmt.Errorf("could not convert range value to int due to %w", err)
	}

	if first > last {
		return ranger{}, fmt.Errorf("first range value is higher than last")
	}

	for i := first; i <= last; i++ {
		if i == first {
			output = fmt.Sprintf("%v", i)
		} else {
			output = fmt.Sprintf("%v %v", output, i)
		}
	}

	return ranger{
		value: output,
	}, nil
}

type ranger struct {
	value string
}

func (i ranger) Value() string {
	return i.value
}
