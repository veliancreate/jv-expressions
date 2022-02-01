package expressions

import (
	"testing"

	"github.com/veliancreate/jv-expressions/parsers"
)

func TestNewCronExpression(t *testing.T) {
	tests := []struct {
		name   string
		arg    string
		expr   *Expression
		parser parsers.CronParser
	}{
		{
			name: "initial example",
			arg:  `* * * * * /usr/bin/find`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := parsers.NewCronParser(tt.arg)
			if err != nil {
				t.Fatalf("error in getting cron parser %v", err)
			}

			_, err = NewCronExpression(parser)
			if err != nil {
				t.Fatalf("error in getting cron expression %v", err)
			}
		})
	}
}
