package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestStepParser(t *testing.T) {
	tests := []struct {
		name        string
		unit        timeunits.MinAndMaxGetter
		shouldError bool
		expected    string
		val         string
	}{
		{
			name:        "step parser prints - day of week 1 step",
			unit:        timeunits.NewDayOfWeek(),
			shouldError: false,
			expected:    "1 2 3 4 5 6 7",
			val:         "1",
		},
		{
			name:        "step parser prints - day of week 2 step",
			unit:        timeunits.NewDayOfWeek(),
			shouldError: false,
			expected:    "1 3 5 7",
			val:         "2",
		},
		{
			name:        "step parser prints - day of week 3 step",
			unit:        timeunits.NewDayOfWeek(),
			shouldError: false,
			expected:    "1 4 7",
			val:         "3",
		},
		{
			name:        "step parser prints - day of month 4 step",
			unit:        timeunits.NewDayOfMonth(),
			shouldError: false,
			expected:    "1 13 25",
			val:         "12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := newStep(tt.val, tt.unit)
			if err != nil && !tt.shouldError {
				t.Fatalf("error in parsing %v", err)
			}

			if parsed.Value() != tt.expected {
				t.Fatalf("expected %v got %v", tt.expected, parsed)
			}
		})
	}
}
