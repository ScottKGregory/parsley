package operations

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		name   string
		a, b   any
		result any
	}{
		{a: 3, b: 3, result: 1},
		{a: 3.4, b: 3, result: 1.1333333333333333},
		{a: 3.4, b: -3, result: -1.1333333333333333},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v/%v", tc.a, tc.b), func(t *testing.T) {
			divide := &DivideOperation{}
			actual, err := divide.Calculate(tc.a, tc.b)
			assert.Nil(t, err)
			assert.Equal(t, tc.result, actual)
		})
	}

}

func TestDivideString(t *testing.T) {
	assert.Equal(t, "/", (&DivideOperation{}).String())
}
