package nodes

import (
	"errors"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestBinaryNode(t *testing.T) {
	testCases := []struct {
		left         *MockNode
		right        *MockNode
		op           *MockBinaryNodeOp
		result       any
		err          error
		stringResult string
	}{
		{
			left:         NewMockNode(nil, 14, nil, "14"),
			right:        NewMockNode(nil, 12, nil, "12"),
			op:           NewMockBinaryNodeOp(12, 14, -12, nil, "-"),
			result:       -12,
			err:          nil,
			stringResult: "14-12",
		},
		{
			left:         NewMockNode(nil, 14, nil, "14"),
			right:        NewMockNode(nil, 0, errors.New("uh oh"), "foobar"),
			op:           NewMockBinaryNodeOp(nil, nil, nil, nil, "-"),
			result:       nil,
			err:          errors.New("rhs error: uh oh"),
			stringResult: "14-foobar",
		},
		{
			left:         NewMockNode(nil, 0, errors.New("uh oh"), "foobar"),
			right:        NewMockNode(nil, 14, nil, "14"),
			op:           NewMockBinaryNodeOp(nil, nil, nil, nil, "-"),
			result:       nil,
			err:          errors.New("lhs error: uh oh"),
			stringResult: "foobar-14",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.stringResult, func(t *testing.T) {
			n := NewBinaryNode(tc.left, tc.right, tc.op)

			res, err := n.Eval(nil)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, res)

			assert.Equal(t, tc.stringResult, n.String())

			tc.left.AssertEvalCalled(t)
			tc.left.AssertStringCalled(t)

			if tc.left.evalErr == nil {
				tc.right.AssertEvalCalled(t)
			}

			tc.right.AssertStringCalled(t)

			tc.op.AssertCalculateCalled(t)
			tc.op.AssertStringCalled(t)
		})
	}
}
