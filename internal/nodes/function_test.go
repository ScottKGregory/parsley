package nodes

import (
	"errors"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestFunctionNode(t *testing.T) {
	// data1 := map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}}

	testCases := []struct {
		name         string
		fun          func(args ...any) (any, error)
		data         map[string]any
		args         []Node
		err          error
		result       any
		stringResult string
	}{
		{
			name: "foo",
			fun: func(args ...any) (any, error) {
				if args[0] != 12 {
					panic("expected 12")
				}
				return 26, nil
			},
			data:         nil,
			args:         []Node{NewMockNode(nil, 12, nil, "12")},
			err:          nil,
			result:       26,
			stringResult: "foo(12)",
		},
		{
			name: "bar",
			fun: func(args ...any) (any, error) {
				if args[0] != 12 {
					panic("expected 12")
				}
				return 26, errors.New("uh oh")
			},
			data:         nil,
			args:         []Node{NewMockNode(nil, 12, nil, "12")},
			err:          errors.New("node evaluation failed: uh oh"),
			result:       nil,
			stringResult: "bar(12)",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n := NewFunctionNode(tc.fun, tc.name, tc.args...)

			res, err := n.Eval(tc.data)
			assert.ErrorEqual(t, tc.err, err)
			assert.Equal(t, tc.result, res)

			assert.Equal(t, tc.stringResult, n.String())
		})
	}
}
