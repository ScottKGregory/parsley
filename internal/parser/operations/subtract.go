package operations

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// SubtractOperation is used to subtract two numbers from each other
type SubtractOperation struct {
}

// Calculate subtracts two numbers from each other
func (o *SubtractOperation) Calculate(a, b any) (any, error) {
	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	return aa - bb, nil
}

// String returns the string representation
func (o *SubtractOperation) String() string { return "-" }
