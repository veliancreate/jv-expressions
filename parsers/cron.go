package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veliancreate/jv-expressions/timeunits"
)

type cronUnitParser struct {
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func newCronUnitParser(minAndMaxGetter timeunits.MinAndMaxGetter) cronUnitParser {
	return cronUnitParser{
		minAndMaxGetter,
	}
}

func (cp cronUnitParser) Parse(val string) (string, error) {
	if val == "*" {
		return newAny(cp.minAndMaxGetter)
	}

	if i, err := strconv.Atoi(val); err == nil {
		return newInteger(i, cp.minAndMaxGetter)
	}

	splitStep := strings.Split(val, "/")
	if len(splitStep) == 2 && splitStep[0] == "*" {
		return newStep(splitStep[1], cp.minAndMaxGetter)
	}

	splitList := strings.Split(val, ",")
	if len(splitList) > 1 {
		return newList(splitList, cp.minAndMaxGetter)
	}

	splitRange := strings.Split(val, "-")
	if len(splitRange) == 2 {
		return newRange([]string{splitRange[0], splitRange[1]}, cp.minAndMaxGetter)
	}

	return "", fmt.Errorf("unsupported type")
}

type CronParser struct {
	Minute     cronUnitParser
	Hour       cronUnitParser
	DayOfMonth cronUnitParser
	Month      cronUnitParser
	DayOfWeek  cronUnitParser
	Expression []string
}

func NewCronParser(rawExpr string) (CronParser, error) {
	expr := strings.Split(rawExpr, " ")

	if len(expr) != 6 {
		return CronParser{}, fmt.Errorf("not enough expressions")
	}

	return CronParser{
		Minute:     newCronUnitParser(timeunits.NewMinute()),
		Hour:       newCronUnitParser(timeunits.NewHour()),
		DayOfMonth: newCronUnitParser(timeunits.NewDayOfMonth()),
		Month:      newCronUnitParser(timeunits.NewMonth()),
		DayOfWeek:  newCronUnitParser(timeunits.NewDayOfWeek()),
		Expression: expr,
	}, nil
}
