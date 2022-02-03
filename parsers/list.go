package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newList(val []string, minAndMaxGetter timeunits.MinAndMaxGetter) (list, error) {
	var output string

	for i, listItem := range val {
		value, err := stringToInt(listItem, minAndMaxGetter)
		if err != nil {
			return list{}, fmt.Errorf("could not convert list value to int due to %w", err)
		}

		if i == 0 {
			output = fmt.Sprintf("%v", value)
		} else {
			output = fmt.Sprintf("%v %v", output, value)
		}
	}

	return list{
		value: output,
	}, nil
}

type list struct {
	value string
}

func (l list) Value() string {
	return l.value
}
