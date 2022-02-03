# jv-expressions

### Usage

Assuming you have the latest version of go installed and the GO111MODULE env variable set to "on"

- Run `go install` from the root
- Run `jv-expressions "<expression>"`

### Tests

- Run `go test ./...`

### Implementation

- The `NewCronParser` function takes a raw expression from the input and creates a `CronParser`
- The `CronParser` has a `cronTimeUnitParser` for each time unit, created by injecting each from the `timeunit` module as well as the value of the input string for that unit
- When the Parse method is called on the `CronParser`, the string input is evaluated and the appropriate expression type is created which implements the `ExpressionTimeUnit` interface
- The value of that process is then available as a `Value` method on the `ExpressionTimeUnit` interface
- The `NewCronExpression` creator takes a `CronParser` and calls `Parse` on each unit parser, then passes the value for which is used on the `Expression` struct
- The expression output can be formatted calling the `ToString` method on the `Expression` struct
