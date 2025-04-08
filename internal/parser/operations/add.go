// Package operations provides implementations or various expression operations
package operations

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// AddOperation is used to add two numbers (or strings) together
type AddOperation struct {
}

// Calculate adds two numbers (or strings) together
func (o *AddOperation) Calculate(a, b any) (any, error) {
	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	return aa + bb, nil
}

// String returns the string representation
func (o *AddOperation) String() string { return "+" }
