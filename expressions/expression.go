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

	minute, err := cronParser.Minute.Parse()
	if err != nil {
		return expression, fmt.Errorf("could not parse minute due to %w", err)
	}

	hour, err := cronParser.Hour.Parse()
	if err != nil {
		return expression, fmt.Errorf("could not parse hour due to %w", err)
	}

	dayOfMonth, err := cronParser.DayOfMonth.Parse()
	if err != nil {
		return expression, fmt.Errorf("could not parse day of month due to %w", err)
	}

	month, err := cronParser.Month.Parse()
	if err != nil {
		return expression, fmt.Errorf("could not parse month due to %w", err)
	}

	dayOfWeek, err := cronParser.DayOfWeek.Parse()
	if err != nil {
		return expression, fmt.Errorf("could not parse day of week due to %w", err)
	}

	return &Expression{
		Minute:     minute.Value(),
		Hour:       hour.Value(),
		DayOfMonth: dayOfMonth.Value(),
		Month:      month.Value(),
		DayOfWeek:  dayOfWeek.Value(),
		Command:    cronParser.Command,
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
