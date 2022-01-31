package parsers

import "github.com/veliancreate/jv-expressions/timeunits"

func NewList(val string, minAndMaxGetter timeunits.MinAndMaxGetter) List {
	return List{
		val,
		minAndMaxGetter,
	}
}

type List struct {
	val             string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s List) Parse() string {
	return ""
}
