package parsers

import (
	"fmt"
	"strconv"

	"github.com/veliancreate/jv-expressions/timeunits"
)

type UnitExpression interface {
	Value() string
}

func stringToInt(item string, getter timeunits.MinAndMaxGetter) (int, error) {
	min := getter.Min()
	max := getter.Max()

	stringAsInt, err := strconv.Atoi(item)
	if err != nil {
		return 0, fmt.Errorf("parsing value as int due to %w", err)
	}

	if stringAsInt > max {
		return 0, fmt.Errorf("value higher than max")
	}

	if stringAsInt < min {
		return 0, fmt.Errorf("value lower than min")
	}

	return stringAsInt, nil
}
