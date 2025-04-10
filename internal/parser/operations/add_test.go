package operations

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name   string
		a, b   any
		result any
	}{
		{a: 3, b: 3, result: 6},
		{a: 3.4, b: 3, result: 6.4},
		{a: 3.4, b: -3, result: 0.3999999999999999},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v+%v", tc.a, tc.b), func(t *testing.T) {
			add := &AddOperation{}
			actual, err := add.Calculate(tc.a, tc.b)
			assert.Nil(t, err)
			assert.Equal(t, tc.result, actual)
		})
	}

}

func TestAddString(t *testing.T) {
	assert.Equal(t, "+", (&AddOperation{}).String())
}
