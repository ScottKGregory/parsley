package operations

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestSubtraction(t *testing.T) {
	testCases := []struct {
		name   string
		a, b   any
		result any
	}{
		{a: 3, b: 3, result: 0},
		{a: 3.4, b: 3, result: 0.3999999999999999},
		{a: 3.4, b: -3, result: 6.4},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.a, tc.b), func(t *testing.T) {
			subtract := &SubtractOperation{}
			actual, err := subtract.Calculate(tc.a, tc.b)
			assert.Nil(t, err)
			assert.Equal(t, tc.result, actual)
		})
	}
}

func TestSubtractionString(t *testing.T) {
	assert.Equal(t, "-", (&SubtractOperation{}).String())
}
