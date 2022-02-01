package parsers

import (
	"fmt"

	"github.com/veliancreate/jv-expressions/timeunits"
)

func NewInteger(val int, minAndMaxGetter timeunits.MinAndMaxGetter) Int {
	return Int{
		val,
		minAndMaxGetter,
	}
}

type Int struct {
	val             int
	minAndMaxGetter timeunits.MinAndMaxGetter
}

func (s Int) Parse() (string, error) {
	max := s.minAndMaxGetter.Max()
	min := s.minAndMaxGetter.Min()

	if s.val > max {
		return "", fmt.Errorf("val %v is over max %v", s.val, max)
	}

	if s.val < min {
		return "", fmt.Errorf("val %v is under min %v", s.val, min)
	}

	return fmt.Sprintf("%v", s.val), nil
}
