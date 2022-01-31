package expressions

import (
	"testing"
)

func TestNewCronExpression(t *testing.T) {
	tests := []struct {
		name   string
		arg    string
		expr   *Expression
		output string
	}{
		{
			name: "initial example",
			arg:  `* * * * * /usr/bin/find`,
			expr: &Expression{
				Minute:     "*",
				Hour:       "*",
				DayOfMonth: "*",
				Month:      "*",
				DayOfWeek:  "*",
				Command:    "/usr/bin/find",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCronExpression(tt.arg)
			if err != nil {
				t.Fatalf("error in getting cron expression %v", err)
			}
		})
	}
}
