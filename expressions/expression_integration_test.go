package expressions

import (
	"testing"

	"github.com/veliancreate/jv-expressions/parsers"
)

func TestNewCronExpression(t *testing.T) {
	tests := []struct {
		name   string
		arg    []string
		expr   *Expression
		parser parsers.CronParser
	}{
		{
			name: "initial example",
			arg:  []string{"1", "2", "3", "4", "5", "/usr/bin/find"},
			expr: &Expression{
				Minute:     "1",
				Hour:       "2",
				DayOfMonth: "3",
				Month:      "4",
				DayOfWeek:  "5",
				Command:    "/usr/bin/find",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := parsers.NewCronParser(tt.arg)
			if err != nil {
				t.Fatalf("error in getting cron parser %v", err)
			}

			e, err := NewCronExpression(parser)
			if err != nil {
				t.Fatalf("error in getting cron expression %v", err)
			}

			if e.Minute != tt.expr.Minute {
				t.Fatalf("expected %v got %v", e.Minute, tt.expr.Minute)
			}

			if e.Hour != tt.expr.Hour {
				t.Fatalf("expected %v got %v", e.Hour, tt.expr.Hour)
			}

			if e.DayOfMonth != tt.expr.DayOfMonth {
				t.Fatalf("expected %v got %v", e.DayOfMonth, tt.expr.DayOfMonth)
			}

			if e.Month != tt.expr.Month {
				t.Fatalf("expected %v got %v", e.Month, tt.expr.Month)
			}

			if e.DayOfMonth != tt.expr.DayOfMonth {
				t.Fatalf("expected %v got %v", e.DayOfMonth, tt.expr.DayOfMonth)
			}

			if e.Command != tt.expr.Command {
				t.Fatalf("expected %v got %v", e.DayOfMonth, tt.expr.DayOfMonth)
			}
		})
	}
}
