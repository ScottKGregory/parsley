package operations

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestNegate(t *testing.T) {
	testCases := []struct {
		name   string
		a      any
		result any
	}{
		{a: 3, result: -3},
		{a: 3.4, result: -3.4},
		{a: -3.4, result: 3.4},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.a), func(t *testing.T) {
			negate := &NegateOperation{}
			actual, err := negate.Calculate(tc.a)
			assert.Nil(t, err)
			assert.Equal(t, tc.result, actual)
		})
	}

}

func TestNegateString(t *testing.T) {
	assert.Equal(t, "-", (&NegateOperation{}).String())
}
