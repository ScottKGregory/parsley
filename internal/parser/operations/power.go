package operations

import (
	"fmt"
	"math"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// PowerOperation is used to calculate the exponent of two numbers
type PowerOperation struct {
}

// Calculate calculates the exponent of two numbers
func (o *PowerOperation) Calculate(a, b any) (any, error) {
	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	return math.Pow(aa, bb), nil
}

// String returns the string representation
func (o *PowerOperation) String() string { return "^" }
