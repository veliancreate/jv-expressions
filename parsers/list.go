package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func newList(val []string, minAndMaxGetter timeunits.MinAndMaxGetter) (string, error) {
	var output string

	for i, listItem := range val {
		value, err := stringToInt(listItem, minAndMaxGetter)
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
