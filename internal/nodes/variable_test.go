package nodes

import (
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestVariableEval(t *testing.T) {
	testCases := []struct {
		a            string
		data         map[string]any
		result       any
		stringResult string
	}{
		{"foo", map[string]any{"foo": "bar"}, "bar", "foo"},
		{"foo", map[string]any{"foo": 2}, 2, "foo"},
		{"foo.bar", map[string]any{"foo": map[string]any{"bar": "baz"}}, "baz", "foo.bar"},
	}
	for _, tc := range testCases {
		t.Run(tc.a, func(t *testing.T) {
			n := NewVariableNode(tc.a)

			res, err := n.Eval(tc.data)
			assert.Equal(t, tc.result, res)
			assert.Equal(t, nil, err)

			assert.Equal(t, tc.stringResult, n.String())
		})
	}
}
