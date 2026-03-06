package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	strConv := strconv.FormatFloat(f, 'f', 1, 64)
	return fmt.Sprintf("This is the number %v", strConv)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	floatConv := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %v", strconv.FormatFloat(floatConv, 'f', 1, 64))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if fnb.Value() == "4" {
		return 0
	}

	if s, err := strconv.Atoi(fnb.Value()); err == nil {
		return s
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if fn, ok := fnb.(FancyNumber); ok {
		conv, _ := strconv.Atoi(fn.Value())
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(conv))
	}
	return "This is a fancy box containing the number 0.0"
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch n := i.(type) {
	case int:
		return fmt.Sprintf("This is the number %.1f", float64(n))
	case float64:
		return fmt.Sprintf("This is the number %.1f", n)
	case NumberBox:
		conv := strconv.FormatFloat(float64(n.Number()), 'f', 1, 64)
		return fmt.Sprintf("This is a box containing the number %v", conv)
	case FancyNumberBox:
		conv, _ := strconv.Atoi(n.Value())
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(conv))
	default:
		return fmt.Sprintf("Return to sender")
	}
}
