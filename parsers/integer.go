package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newInteger(val int, minAndMaxGetter timeunits.MinAndMaxGetter) (string, error) {
	max := minAndMaxGetter.Max()
	min := minAndMaxGetter.Min()

	if val > max {
		return "", fmt.Errorf("val %v is over max %v", val, max)
	}

	if val < min {
		return "", fmt.Errorf("val %v is under min %v", val, min)
	}

	return fmt.Sprintf("%v", val), nil
}
