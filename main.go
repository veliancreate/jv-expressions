package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/veliancreate/jv-expressions/expressions"
	"github.com/veliancreate/jv-expressions/parsers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("not enough args provided")
		os.Exit(1)
	}

	rawExpr := os.Args[1]

	expr := strings.Split(rawExpr, " ")

	parser, err := parsers.NewCronParser(expr)
	if err != nil {
		fmt.Printf("getting new cron parser :: %v", err)
		os.Exit(1)
	}

	expression, err := expressions.NewCronExpression(parser)
	if err != nil {
		fmt.Printf("getting new cron expression :: %v", err)
		os.Exit(1)
	}

	fmt.Println(expression.ToString())
}
