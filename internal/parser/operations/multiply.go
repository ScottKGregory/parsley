package operations

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// MultiplyOperation is used to multiply two numbers
type MultiplyOperation struct {
}

// Calculate multiplies two numbers
func (o *MultiplyOperation) Calculate(a, b any) (any, error) {
	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	return aa * bb, nil
}

// String returns the string representation
func (o *MultiplyOperation) String() string { return "*" }
