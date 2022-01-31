package parsers

import "github.com/veliancreate/jv-expressions/timeunits"

func NewRange(val string, minAndMaxGetter timeunits.MinAndMaxGetter) Range {
	return Range{
		val,
		minAndMaxGetter,
	}
}

type Range struct {
	val             string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s Range) Parse() string {
	return ""
}
