package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestRangeParser(t *testing.T) {
	tests := []struct {
		name        string
		shouldError bool
		expected    string
		unit        timeunits.MinAndMaxGetter
		val         []string
	}{
		{
			name:        "list parser prints",
			shouldError: false,
			unit:        timeunits.NewDayOfWeek(),
			val:         []string{"1", "5"},
			expected:    "1 2 3 4 5",
		},
		{
			name:        "range parser value lower than min",
			unit:        timeunits.NewDayOfWeek(),
			val:         []string{"0", "2"},
			shouldError: true,
		},
		{
			name:        "range parser value higher than max",
			unit:        timeunits.NewDayOfWeek(),
			val:         []string{"8", "2"},
			shouldError: true,
		},
		{
			name:        "first value higher than last",
			unit:        timeunits.NewDayOfWeek(),
			val:         []string{"5", "2"},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := newRange(tt.val, tt.unit)
			if err != nil && !tt.shouldError {
				t.Fatalf("error in parsing %v", err)
			}

			if parsed.Value() != tt.expected {
				t.Fatalf("expected %v got %v", tt.expected, parsed)
			}
		})
	}
}
