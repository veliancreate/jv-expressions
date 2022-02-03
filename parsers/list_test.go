package parsers

import (
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestListParser(t *testing.T) {
	tests := []struct {
		name        string
		shouldError bool
		expected    string
		unit        timeunits.Unit
		val         []string
	}{
		{
			name:        "list parser prints",
			shouldError: false,
			expected:    "1 2 3",
			val:         []string{"1", "2", "3"},
			unit:        timeunits.NewDayOfWeek(),
		},
		{
			name:        "list parser value lower than min",
			shouldError: true,
			val:         []string{"0", "2", "3"},
			unit:        timeunits.NewDayOfWeek(),
		},
		{
			name:        "list parser value higher than max",
			shouldError: true,
			val:         []string{"8", "2", "3"},
			unit:        timeunits.NewDayOfWeek(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := newList(tt.val, tt.unit)
			if err != nil && !tt.shouldError {
				t.Fatalf("error in parsing %v", err)
			}

			if parsed != tt.expected {
				t.Fatalf("expected %v got %v", tt.expected, parsed)
			}
		})
	}
}
