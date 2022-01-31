package expressions

import (
	"fmt"
	"strings"

	"github.com/veliancreate/jv-expressions/parsers"
	"github.com/veliancreate/jv-expressions/timeunits"
)

type Expression struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

func validateExpression(expression []string) error {
	const standardValidCharacters = `^[,*\d\/-]+$`
	const dayOfWeekValidCharacters = `^[?,*\dL#/-]+$`
	const dayOfMonthValidCharacters = `^[?,*\dL/-]+$`
	return nil
}

func getType(val string, getter timeunits.MinAndMaxGetter) (parsers.Parser, error) {
	if val == "*" {
		return parsers.NewAny(val, getter), nil
	}

	// if _, err := strconv.Atoi(val); err == nil {
	// 	return parsers.NewInt(val, getter), nil
	// }

	// split := strings.Split(val, "/")
	// if len(split) == 2 && split[0] == "*" {
	// 	return parsers.NewStep(val, getter), nil
	// }

	// split = strings.Split(val, ",")
	// if len(split) > 1 {
	// 	return parsers.NewList(val, getter), nil
	// }

	// split = strings.Split(val, "-")
	// if len(split) == 2 {
	// 	return parsers.NewRange(val, getter), nil
	// }

	return nil, fmt.Errorf("unsupported type")
}

func NewCronExpression(rawExpression string) (*Expression, error) {
	var expression *Expression

	expressions := strings.Split(rawExpression, " ")
	if len(expressions) != 6 {
		return expression, fmt.Errorf("not enough args in expression")
	}

	err := validateExpression(expressions)
	if err != nil {
		return expression, fmt.Errorf("validating expression")
	}

	minute, err := getType(expressions[0], timeunits.NewMinute())
	if err != nil {
		return expression, fmt.Errorf("getting minute type :: %v", err)
	}

	hour, err := getType(expressions[1], timeunits.NewHour())
	if err != nil {
		return expression, fmt.Errorf("getting hour type :: %v", err)
	}

	dayOfMonth, err := getType(expressions[2], timeunits.NewDayOfMonth())
	if err != nil {
		return expression, fmt.Errorf("getting day of month type :: %v", err)
	}

	month, err := getType(expressions[3], timeunits.NewMonth())
	if err != nil {
		return expression, fmt.Errorf("getting month type :: %v", err)
	}

	dayOfWeek, err := getType(expressions[4], timeunits.NewDayOfWeek())
	if err != nil {
		return expression, fmt.Errorf("getting dayOfWeek type :: %v", err)
	}

	return &Expression{
		Minute:     minute.Parse(),
		Hour:       hour.Parse(),
		DayOfMonth: dayOfMonth.Parse(),
		Month:      month.Parse(),
		DayOfWeek:  dayOfWeek.Parse(),
		Command:    expressions[5],
	}, nil
}

func (e *Expression) ToString() string {
	return fmt.Sprintf(`
		minute 			%s
		hour			%s
		day of month 	%s
		month 			%s
		day of week 	%s
		command 		%s  
	`,
		e.Minute,
		e.Hour,
		e.DayOfMonth,
		e.Month,
		e.DayOfWeek,
		e.Command,
	)
}
