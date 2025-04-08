package operations

import (
	"github.com/scottkgregory/parsley/internal/helpers"
)

// NegateOperation is used to negate a number
type NegateOperation struct {
}

// Calculate negates a number
func (o *NegateOperation) Calculate(a any) (any, error) {
	aa, err := helpers.ToFloat64(a)
	return -aa, err
}

// String returns the string representation
func (o *NegateOperation) String() string { return "-" }
