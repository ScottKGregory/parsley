package nodes

import (
	"fmt"
)

// UnaryNodeOp is the interface for an operation that works on a unary node
type UnaryNodeOp interface {
	NodeOp
	Calculate(right any) (any, error)
}

// UnaryNode is a node that has just a right side
type UnaryNode struct {
	Right Node
	op    UnaryNodeOp
}

// NewUnaryNode creates a nwe unary node
func NewUnaryNode(right Node, op UnaryNodeOp) *UnaryNode {
	return &UnaryNode{right, op}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *UnaryNode) Eval(data map[string]any) (any, error) {
	val, err := n.Right.Eval(data)
	if err != nil {
		return nil, err
	}

	return n.op.Calculate(val)
}

// String returns the string representation
func (n *UnaryNode) String() string {
	return fmt.Sprintf("%s(%s)", n.op.String(), n.Right.String())
}
