package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestRangeParser(t *testing.T) {
	tests := []struct {
		name        string
		parser      Parser
		shouldError bool
		expected    string
	}{
		{
			name:        "list parser prints",
			parser:      NewRange([]string{"1", "5"}, timeunits.NewDayOfWeek()),
			shouldError: false,
			expected:    "1 2 3 4 5",
		},
		{
			name:        "range parser value lower than min",
			parser:      NewRange([]string{"0", "2"}, timeunits.NewDayOfWeek()),
			shouldError: true,
		},
		{
			name:        "range parser value higher than max",
			parser:      NewRange([]string{"8", "2"}, timeunits.NewDayOfWeek()),
			shouldError: true,
		},
		{
			name:        "first value higher than last",
			parser:      NewRange([]string{"5", "2"}, timeunits.NewDayOfWeek()),
			shouldError: true,
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
