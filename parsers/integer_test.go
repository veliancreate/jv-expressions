package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestNumberParser(t *testing.T) {
	tests := []struct {
		name        string
		val         int
		shouldError bool
		expected    string
		unit        timeunits.Unit
	}{
		{
			name:        "fails when the input is below the min",
			val:         0,
			unit:        timeunits.NewDayOfWeek(),
			shouldError: true,
		},
		{
			name:        "fails when the input is below the max",
			val:         8,
			unit:        timeunits.NewDayOfWeek(),
			shouldError: true,
		},
		{
			name:        "returns the integer",
			val:         7,
			unit:        timeunits.NewDayOfWeek(),
			shouldError: false,
			expected:    "7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := newInteger(tt.val, tt.unit)
			if err != nil && !tt.shouldError {
				t.Fatalf("error in parsing %v", err)
			}

			if parsed != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, parsed)
			}
		})
	}
}
