package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestAnyParser(t *testing.T) {
	tests := []struct {
		name     string
		parser   Parser
		expected string
	}{
		{
			name:     "prints out the values between the min and max - day of week",
			parser:   NewAny("", timeunits.NewDayOfWeek()),
			expected: "1 2 3 4 5 6 7",
		},
		{
			name:     "prints out the values between the min and max - month",
			parser:   NewAny("", timeunits.NewMonth()),
			expected: "1 2 3 4 5 6 7 8 9 10 11 12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := tt.parser.Parse()
			if err != nil {
				t.Fatalf("error parsing * %v", err)
			}

			if parsed != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, parsed)
			}
		})
	}
}
