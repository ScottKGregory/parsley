package nodes

import (
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestStringEval(t *testing.T) {
	testCases := []struct {
		a            string
		result       string
		stringResult string
	}{
		{"foo", "foo", `"foo"`},
		{"asidfughasuhf", "asidfughasuhf", `"asidfughasuhf"`},
		{"foo\"bar", "foo\"bar", `"foo\"bar"`},
		{"foo", "foo", `"foo"`},
	}
	for _, tc := range testCases {
		t.Run(tc.a, func(t *testing.T) {
			n := NewStringNode(tc.a)

			res, err := n.Eval(nil)
			assert.Equal(t, tc.result, res)
			assert.Equal(t, nil, err)

			assert.Equal(t, tc.stringResult, n.String())
		})
	}
}
