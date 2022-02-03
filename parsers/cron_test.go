package parsers

import (
	"reflect"
	"testing"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func TestCronUnitParser(t *testing.T) {
	tests := []struct {
		name     string
		expected reflect.Type
		value    string
	}{
		{
			name:     "cron parser creates any",
			value:    "*",
			expected: reflect.TypeOf(any{}),
		},
		{
			name:     "cron parser creates integer",
			value:    "1",
			expected: reflect.TypeOf(integer{}),
		},
		{
			name:     "cron parser creates list",
			value:    "1,2",
			expected: reflect.TypeOf(list{}),
		},
		{
			name:     "cron parser creates ranger",
			value:    "1-2",
			expected: reflect.TypeOf(ranger{}),
		},
		{
			name:     "cron parser creates step",
			value:    "*/2",
			expected: reflect.TypeOf(step{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := newCronUnitParser(timeunits.NewDayOfWeek(), tt.value).Parse()
			if err != nil {
				t.Fatalf("unexpected error in parsing %v", err)
			}

			parsedType := reflect.TypeOf(parsed)

			if parsedType.Name() != tt.expected.Name() {
				t.Fatalf("parsed type incorrect, expected %v got %v", tt.expected.Name(), parsedType.Name())
			}
		})
	}
}

func TestCronUnitParserValidation(t *testing.T) {
	tests := []struct {
		name        string
		shouldError bool
		unit        timeunits.Unit
		value       string
	}{
		{
			name:  "cron parser validation",
			value: "/1/5",
			unit:  timeunits.NewDayOfWeek(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newCronUnitParser(timeunits.NewDayOfWeek(), tt.value).Parse()
			if err == nil {
				t.Fatalf("this should fail validation check")
			}
		})
	}
}
