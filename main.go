package main

import (
	"fmt"
	"os"

	"github.com/veliancreate/jv-expressions/expressions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("not enough args provided")
		os.Exit(1)
	}

	rawExpr := os.Args[1]

	expression, err := expressions.NewCronExpression(rawExpr)
	if err != nil {
		fmt.Printf("getting new cron expression :: %v", err)
		os.Exit(1)
	}

	fmt.Println(expression.ToString())
}
