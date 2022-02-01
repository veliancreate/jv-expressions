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
		return NewRange([]string{splitRange[0], splitRange[2]}, getter), nil
	}

	return nil, fmt.Errorf("unsupported type")
}

func NewCronParser(expressions []string) (CronParser, error) {
	var cronParser CronParser

	minuteParser, err := getParser(expressions[0], timeunits.NewMinute())
	if err != nil {
		return cronParser, fmt.Errorf("getting minute type :: %w", err)
	}

	hourParser, err := getParser(expressions[1], timeunits.NewHour())
	if err != nil {
		return cronParser, fmt.Errorf("getting hour type :: %w", err)
	}

	dayOfMonthParser, err := getParser(expressions[2], timeunits.NewDayOfMonth())
	if err != nil {
		return cronParser, fmt.Errorf("getting day of month type :: %w", err)
	}

	monthParser, err := getParser(expressions[3], timeunits.NewMonth())
	if err != nil {
		return cronParser, fmt.Errorf("getting month type :: %w", err)
	}

	dayOfWeekParser, err := getParser(expressions[4], timeunits.NewDayOfWeek())
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
