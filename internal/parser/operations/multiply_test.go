package operations

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name   string
		a, b   any
		result any
	}{
		{a: 3, b: 3, result: 9},
		{a: 3.4, b: 3, result: 10.2},
		{a: 3.4, b: -3, result: -10.2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v*%v", tc.a, tc.b), func(t *testing.T) {
			multiply := &MultiplyOperation{}
			actual, err := multiply.Calculate(tc.a, tc.b)
			assert.Nil(t, err)
			assert.Equal(t, tc.result, actual)
		})
	}

}

func TestMultiplyString(t *testing.T) {
	assert.Equal(t, "*", (&MultiplyOperation{}).String())
}
