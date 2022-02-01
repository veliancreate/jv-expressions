package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestStepParser(t *testing.T) {
	tests := []struct {
		name        string
		parser      Parser
		shouldError bool
		expected    string
	}{
		{
			name:        "step parser prints - day of week 1 step",
			parser:      NewStep("1", timeunits.NewDayOfWeek()),
			shouldError: false,
			expected:    "1 2 3 4 5 6 7",
		},
		{
			name:        "step parser prints - day of week 2 step",
			parser:      NewStep("2", timeunits.NewDayOfWeek()),
			shouldError: false,
			expected:    "1 3 5 7",
		},
		{
			name:        "step parser prints - day of week 3 step",
			parser:      NewStep("3", timeunits.NewDayOfWeek()),
			shouldError: false,
			expected:    "1 4 7",
		},
		{
			name:        "step parser prints - day of month 4 step",
			parser:      NewStep("12", timeunits.NewDayOfMonth()),
			shouldError: false,
			expected:    "1 13 25",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := tt.parser.Parse()
			if err != nil && !tt.shouldError {
				t.Fatalf("error in parsing %v", err)
			}

			if parsed != tt.expected {
				t.Fatalf("expected %v got %v", tt.expected, parsed)
			}
		})
	}
}
