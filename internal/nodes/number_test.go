package nodes

import (
	"errors"
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestNumberEval(t *testing.T) {
	testCases := []struct {
		a            any
		result       any
		stringResult string
		err          error
	}{
		{12, float64(12), `12`, nil},
		{12.0, float64(12), `12`, nil},
		{12.011, 12.011, `12.011`, nil},
		{"12.011", 12.011, `12.011`, nil},
		{"blep", nil, `blep`, errors.New("node evaluation failed: error parsing value as float, could not parse string 'blep'")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.a), func(t *testing.T) {
			n := NewNumberNode(tc.a)

			res, err := n.Eval(nil)
			assert.Equal(t, tc.result, res)
			assert.ErrorEqual(t, tc.err, err)

			assert.Equal(t, tc.stringResult, n.String())
		})
	}
}
