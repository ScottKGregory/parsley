package nodes

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// UnaryNode is a node that has just a right side
type UnaryNode struct {
	Right Node
	op    string
}

// NewUnaryNode creates a nwe unary node
func NewUnaryNode(right Node, op string) *UnaryNode {
	return &UnaryNode{right, op}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *UnaryNode) Eval(data map[string]any) (any, error) {
	if n.op != "-" {
		return nil, fmt.Errorf("unrecognised op: %s", string(n.op))
	}

	val, err := n.Right.Eval(data)
	if err != nil {
		return nil, err
	}

	aa, err := helpers.ToFloat64(val)
	return -aa, err
}

// String returns the string representation
func (n *UnaryNode) String() string {
	return fmt.Sprintf("%s(%s)", n.op, n.Right.String())
}
