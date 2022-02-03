package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newAny(minAndMaxGetter timeunits.MinAndMaxGetter) (string, error) {
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

	return output, nil
}
