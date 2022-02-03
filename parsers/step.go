package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newStep(val string, minAndMaxGetter timeunits.MinAndMaxGetter) (string, error) {
	var output string
	min := minAndMaxGetter.Min()
	max := minAndMaxGetter.Max()

	value, err := stringToInt(val, minAndMaxGetter)
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
