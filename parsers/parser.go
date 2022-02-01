package parsers

import (
	"github.com/veliancreate/jv-expressions/timeunits"
)

type Parser interface {
	Parse() (string, error)
}

type ParserCreator func(val string, minAndMaxGetter timeunits.MinAndMaxGetter) Parser
