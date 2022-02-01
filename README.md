# jv-expressions

### Usage

Assuming you have the latest version of go installed and the GO111MODULE env variable set to "on"

- Run `go install` from the root
- Run `jv-expressions "<expression>"`

### Tests

- Run `go test ./...`

### Notes

Only the `*` (any) and `<integer>` formats are currently supported. The project scaffolds the other parsers so that the parsing logic itself can be added to the `Parse` function of each type in the `parsers` package, and extending the supported types in the `NewCronParser` function

