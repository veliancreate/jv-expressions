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

func NewCronExpression(parser parsers.CronParser) (*Expression, error) {
	var expression *Expression

	minute, err := parser.Minute.Parse()
	if err != nil {
		return expression, fmt.Errorf("couldnt parse minute :: %w", err)
	}

	hour, err := parser.Hour.Parse()
	if err != nil {
		return expression, fmt.Errorf("couldnt parse hour :: %w", err)
	}

	dayOfMonth, err := parser.DayOfMonth.Parse()
	if err != nil {
		return expression, fmt.Errorf("couldnt parse day of month :: %w", err)
	}

	month, err := parser.Month.Parse()
	if err != nil {
		return expression, fmt.Errorf("couldnt parse month :: %w", err)
	}

	dayOfWeek, err := parser.DayOfWeek.Parse()
	if err != nil {
		return expression, fmt.Errorf("couldnt parse day of week :: %w", err)
	}

	return &Expression{
		Minute:     minute,
		Hour:       hour,
		DayOfMonth: dayOfMonth,
		Month:      month,
		DayOfWeek:  dayOfWeek,
		Command:    parser.Command,
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
