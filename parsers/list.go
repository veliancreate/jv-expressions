package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func NewList(val []string, minAndMaxGetter timeunits.MinAndMaxGetter) List {
	return List{
		val,
		minAndMaxGetter,
	}
}

type List struct {
	val             []string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (l List) Parse() (string, error) {
	var output string

	for i, listItem := range l.val {
		value, err := stringToInt(listItem, l.minAndMaxGetter)
		if err != nil {
			return "", fmt.Errorf("could not convert list value to int due to %w", err)
		}

		if i == 0 {
			output = fmt.Sprintf("%v", value)
		} else {
			output = fmt.Sprintf("%v %v", output, value)
		}
	}

	return output, nil
}
