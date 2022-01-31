package parsers

import "github.com/veliancreate/jv-expressions/timeunits"

func NewStep(val string, minAndMaxGetter timeunits.MinAndMaxGetter) Step {
	return Step{
		val,
		minAndMaxGetter,
	}
}

type Step struct {
	val             string
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s Step) Parse() string {
	return ""
}
