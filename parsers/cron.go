package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veliancreate/jv-expressions/timeunits"
)

type cronUnitParser struct {
	minAndMaxGetter timeunits.MinAndMaxGetter
	val             string
}

func newCronUnitParser(minAndMaxGetter timeunits.MinAndMaxGetter, val string) cronUnitParser {
	return cronUnitParser{
		minAndMaxGetter,
		val,
	}
}

func (cp cronUnitParser) Parse() (UnitExpression, error) {
	val := cp.val

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

	return any{}, fmt.Errorf("unsupported type")
}

type CronParser struct {
	Minute         cronUnitParser
	Hour           cronUnitParser
	DayOfMonth     cronUnitParser
	Month          cronUnitParser
	DayOfWeek      cronUnitParser
	Command        string
	RawExpressions []string
}

func NewCronParser(rawExpr string) (CronParser, error) {
	expr := strings.Split(rawExpr, " ")

	if len(expr) != 6 {
		return CronParser{}, fmt.Errorf("not enough expressions")
	}

	rawMinute := expr[0]
	rawHour := expr[1]
	rawDayOfMonth := expr[2]
	rawMonth := expr[3]
	rawDayOfWeek := expr[4]
	rawCommand := expr[5]

	return CronParser{
		Minute:         newCronUnitParser(timeunits.NewMinute(), rawMinute),
		Hour:           newCronUnitParser(timeunits.NewHour(), rawHour),
		DayOfMonth:     newCronUnitParser(timeunits.NewDayOfMonth(), rawDayOfMonth),
		Month:          newCronUnitParser(timeunits.NewMonth(), rawMonth),
		DayOfWeek:      newCronUnitParser(timeunits.NewDayOfWeek(), rawDayOfWeek),
		Command:        rawCommand,
		RawExpressions: expr,
	}, nil
}
