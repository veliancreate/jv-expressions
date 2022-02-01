package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestListParser(t *testing.T) {
	tests := []struct {
		name        string
		parser      Parser
		shouldError bool
		expected    string
	}{
		{
			name:        "list parser prints",
			parser:      NewList([]string{"1", "2", "3"}, timeunits.NewDayOfWeek()),
			shouldError: false,
			expected:    "1 2 3",
		},
		{
			name:        "list parser value lower than min",
			parser:      NewList([]string{"0", "2", "3"}, timeunits.NewDayOfWeek()),
			shouldError: true,
		},
		{
			name:        "list parser value higher than max",
			parser:      NewList([]string{"8", "2", "3"}, timeunits.NewDayOfWeek()),
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
