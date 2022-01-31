package parsers

import "github.com/veliancreate/jv-expressions/timeunits"

func NewInt(val string, minAndMaxGetter timeunits.MinAndMaxGetter) Int {
	return Int{
		val,
		minAndMaxGetter,
	}
}

type Int struct {
	val             string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s Int) Parse() string {
	return ""
}
