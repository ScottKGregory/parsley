package operations

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// DivideOperation is used to divide two numbers
type DivideOperation struct {
}

// Calculate divides two numbers
func (o *DivideOperation) Calculate(a, b any) (any, error) {
	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	return aa / bb, nil
}

// String returns the string representation
func (o *DivideOperation) String() string { return "/" }
