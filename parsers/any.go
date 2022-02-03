package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newAny(minAndMaxGetter timeunits.MinAndMaxGetter) (any, error) {
	min := minAndMaxGetter.Min()
	max := minAndMaxGetter.Max()
	var output string
	for i := min; i <= max; i++ {
		if i == min {
			output = fmt.Sprintf("%v", i)
		} else {
			output = fmt.Sprintf("%v %v", output, i)
		}
	}

	return any{
		value: output,
	}, nil
}

type any struct {
	value string
}

func (a any) Value() string {
	return a.value
}
