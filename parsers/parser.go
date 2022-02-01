package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veliancreate/jv-expressions/timeunits"
)

type Parser interface {
	Parse() (string, error)
}

func getParser(val string, getter timeunits.MinAndMaxGetter) (Parser, error) {
	if val == "*" {
		return NewAny(val, getter), nil
	}

	if i, err := strconv.Atoi(val); err == nil {
		return NewInteger(i, getter), nil
	}

	splitStep := strings.Split(val, "/")
	if len(splitStep) == 2 && splitStep[0] == "*" {
		return NewStep(splitStep[1], getter), nil
	}

	splitList := strings.Split(val, ",")
	if len(splitList) > 1 {
		return NewList(splitList, getter), nil
	}

	splitRange := strings.Split(val, "-")
	if len(splitRange) == 2 {
		return NewRange([]string{splitRange[0], splitRange[1]}, getter), nil
	}

	return nil, fmt.Errorf("unsupported type")
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
