package expressions

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/parsers"
)

type Expression struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

func NewCronExpression(cronParser parsers.CronParser) (*Expression, error) {
	var expression *Expression

	expr := cronParser.Expression

	rawMinute := expr[0]
	minute, err := cronParser.Minute.Parse(rawMinute)
	if err != nil {
		return expression, fmt.Errorf("could not parse minute due to %w", err)
	}

	rawHour := expr[1]
	hour, err := cronParser.Hour.Parse(rawHour)
	if err != nil {
		return expression, fmt.Errorf("could not parse hour due to %w", err)
	}

	rawDayOfMonth := expr[2]
	dayOfMonth, err := cronParser.DayOfMonth.Parse(rawDayOfMonth)
	if err != nil {
		return expression, fmt.Errorf("could not parse day of month due to %w", err)
	}

	rawMonth := expr[3]
	month, err := cronParser.Month.Parse(rawMonth)
	if err != nil {
		return expression, fmt.Errorf("could not parse month due to %w", err)
	}

	rawDayOfWeek := expr[4]
	dayOfWeek, err := cronParser.DayOfWeek.Parse(rawDayOfWeek)
	if err != nil {
		return expression, fmt.Errorf("could not parse day of week due to %w", err)
	}

	rawCommand := expr[5]

	return &Expression{
		Minute:     minute,
		Hour:       hour,
		DayOfMonth: dayOfMonth,
		Month:      month,
		DayOfWeek:  dayOfWeek,
		Command:    rawCommand,
	}, nil
}

func (e *Expression) ToString() string {
	return fmt.Sprintf(`
		minute: %s
		hour: %s
		day of month: %s
		month: %s
		day of week: %s
		command: %s
	`,
		e.Minute,
		e.Hour,
		e.DayOfMonth,
		e.Month,
		e.DayOfWeek,
		e.Command,
	)
}
