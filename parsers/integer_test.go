package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestNumberParser(t *testing.T) {
	tests := []struct {
		name        string
		parser      Parser
		shouldError bool
		expected    string
	}{
		{
			name:        "fails when the input is below the min",
			parser:      NewInteger(0, timeunits.NewDayOfWeek()),
			shouldError: true,
		},
		{
			name:        "fails when the input is below the max",
			parser:      NewInteger(8, timeunits.NewDayOfWeek()),
			shouldError: true,
		},
		{
			name:        "returns the integer",
			parser:      NewInteger(7, timeunits.NewDayOfWeek()),
			shouldError: false,
			expected:    "7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := tt.parser.Parse()
			if err != nil && !tt.shouldError {
				t.Fatalf("error in parsing %v", err)
			}

			if parsed != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, parsed)
			}
		})
	}
}
