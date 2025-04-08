package operations

import (
	"fmt"
	"strings"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// ComparisonOperation is used to establish equality between two values
type ComparisonOperation struct {
	Comparator string
}

// Calculate establishes equality between two values
func (o *ComparisonOperation) Calculate(a, b any) (any, error) {

	if o.Comparator == "||" || o.Comparator == "&&" {
		x, err := helpers.ToBool(a)
		if err != nil {
			return nil, fmt.Errorf("could not parse as boolean: %T %T", a, b)
		}

		y, err := helpers.ToBool(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse as boolean: %T %T", a, b)
		}

		if o.Comparator == "||" {
			return x || y, nil
		}

		return x && y, nil
	}

	x, aOk := a.(string)
	y, bOk := b.(string)
	if aOk && bOk {
		switch o.Comparator {
		case "<":
			return strings.Compare(x, y) < 0, nil
		case ">":
			return strings.Compare(x, y) > 0, nil
		case "=", "==":
			return x == y, nil
		}
	}

	if aOk || bOk {
		return nil, fmt.Errorf("only one side of comparison was a string: %T %T", a, b)
	}

	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	switch o.Comparator {
	case "<":
		return aa < bb, nil
	case ">":
		return aa > bb, nil
	case "=", "==":
		return aa == bb, nil
	}

	return nil, fmt.Errorf("unrecognised comparator: %s", string(o.Comparator))
}

// String returns the string representation
func (o *ComparisonOperation) String() string {
	return fmt.Sprintf(" %s ", string(o.Comparator))
}
