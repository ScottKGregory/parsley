package nodes

import (
	"errors"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestUnaryNode(t *testing.T) {
	testCases := []struct {
		right        *MockNode
		op           *MockUnaryNodeOp
		result       any
		err          error
		stringResult string
	}{
		{
			right:        NewMockNode(nil, 12, nil, "12"),
			op:           NewMockUnaryNodeOp(12, -12, nil, "-"),
			result:       -12,
			err:          nil,
			stringResult: "-(12)",
		},
		{
			right:        NewMockNode(nil, 0, errors.New("uh oh"), "foobar"),
			op:           NewMockUnaryNodeOp(nil, nil, nil, "-"),
			result:       nil,
			err:          errors.New("uh oh"),
			stringResult: "-(foobar)",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.stringResult, func(t *testing.T) {
			n := NewUnaryNode(tc.right, tc.op)

			res, err := n.Eval(nil)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, res)

			assert.Equal(t, tc.stringResult, n.String())

			tc.right.AssertEvalCalled(t)
			tc.right.AssertStringCalled(t)
			tc.op.AssertCalculateCalled(t)
			tc.op.AssertStringCalled(t)
		})
	}
}
