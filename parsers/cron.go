package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/veliancreate/jv-expressions/timeunits"
)

type CronParser struct {
	Minute     Parser
	Hour       Parser
	DayOfMonth Parser
	Month      Parser
	DayOfWeek  Parser
	Command    string
}

func getType(val string, getter timeunits.MinAndMaxGetter) (Parser, error) {
	if val == "*" {
		return NewAny(val, getter), nil
	}

	if i, err := strconv.Atoi(val); err == nil {
		return NewInteger(i, getter), nil
	}

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

func NewCronParser(rawExpr string) (CronParser, error) {
	var cronParser CronParser

	expressions := strings.Split(rawExpr, " ")
	if len(expressions) != 6 {
		return cronParser, fmt.Errorf("not enough args in expression")
	}

	minuteParser, err := getType(expressions[0], timeunits.NewMinute())
	if err != nil {
		return cronParser, fmt.Errorf("getting minute type :: %w", err)
	}

	hourParser, err := getType(expressions[1], timeunits.NewHour())
	if err != nil {
		return cronParser, fmt.Errorf("getting hour type :: %w", err)
	}

	dayOfMonthParser, err := getType(expressions[2], timeunits.NewDayOfMonth())
	if err != nil {
		return cronParser, fmt.Errorf("getting day of month type :: %w", err)
	}

	monthParser, err := getType(expressions[3], timeunits.NewMonth())
	if err != nil {
		return cronParser, fmt.Errorf("getting month type :: %w", err)
	}

	dayOfWeekParser, err := getType(expressions[4], timeunits.NewDayOfWeek())
	if err != nil {
		return cronParser, fmt.Errorf("getting dayOfWeek type :: %w", err)
	}

	return CronParser{
		Minute:     minuteParser,
		Hour:       hourParser,
		DayOfMonth: dayOfMonthParser,
		Month:      monthParser,
		DayOfWeek:  dayOfWeekParser,
		Command:    expressions[5],
	}, nil
}
