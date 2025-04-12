package nodes

import (
	"fmt"
)

// BinaryNodeOp is the interface for an operation that works on a binary node
type BinaryNodeOp interface {
	NodeOp
	Calculate(left, right any) (any, error)
}

// BinaryNode is a node that has both a left and right side
type BinaryNode struct {
	Left  Node
	Right Node
	op    BinaryNodeOp
}

var _ Node = &BinaryNode{}

// NewBinaryNode creates a new binary node
func NewBinaryNode(left, right Node, op BinaryNodeOp) *BinaryNode {
	return &BinaryNode{left, right, op}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *BinaryNode) Eval(data map[string]any) (any, error) {
	// Evaluate both sides
	leftVal, leftErr := n.Left.Eval(data)
	if leftErr != nil {
		return nil, fmt.Errorf("lhs error: %w", leftErr)
	}
	rightVal, rightErr := n.Right.Eval(data)
	if rightErr != nil {
		return nil, fmt.Errorf("rhs error: %w", rightErr)
	}

	// Evaluate and return
	return n.op.Calculate(leftVal, rightVal)
}

// String returns the string representation
func (n *BinaryNode) String() string {
	return fmt.Sprintf("%s%s%s", n.Left.String(), n.op.String(), n.Right.String())
}
