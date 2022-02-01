package parsers

import (
	"fmt"

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

func NewCronParser(expr []string) (CronParser, error) {
	var cronParser CronParser

	if len(expr) != 6 {
		return cronParser, fmt.Errorf("not enough expressions")
	}

	minute := expr[0]
	hour := expr[1]
	dayOfMonth := expr[2]
	month := expr[3]
	dayOfWeek := expr[4]
	command := expr[5]

	minuteParser, err := getParser(minute, timeunits.NewMinute())
	if err != nil {
		return cronParser, fmt.Errorf("getting minute type :: %w", err)
	}

	hourParser, err := getParser(hour, timeunits.NewHour())
	if err != nil {
		return cronParser, fmt.Errorf("getting hour type :: %w", err)
	}

	dayOfMonthParser, err := getParser(dayOfMonth, timeunits.NewDayOfMonth())
	if err != nil {
		return cronParser, fmt.Errorf("getting day of month type :: %w", err)
	}

	monthParser, err := getParser(month, timeunits.NewMonth())
	if err != nil {
		return cronParser, fmt.Errorf("getting month type :: %w", err)
	}

	dayOfWeekParser, err := getParser(dayOfWeek, timeunits.NewDayOfWeek())
	if err != nil {
		return cronParser, fmt.Errorf("getting dayOfWeek type :: %w", err)
	}

	return CronParser{
		Minute:     minuteParser,
		Hour:       hourParser,
		DayOfMonth: dayOfMonthParser,
		Month:      monthParser,
		DayOfWeek:  dayOfWeekParser,
		Command:    command,
	}, nil
}
