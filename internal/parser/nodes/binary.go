//go:generate go tool ridicule -header -in ./binary.go

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

	result any
}

var _ Node = &BinaryNode{}

// NewBinaryNode creates a new binary node
func NewBinaryNode(left, right Node, op BinaryNodeOp) *BinaryNode {
	return &BinaryNode{left, right, op, 0}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *BinaryNode) Eval() (any, error) {
	// Evaluate both sides
	leftVal, leftErr := n.Left.Eval()
	if leftErr != nil {
		return nil, fmt.Errorf("lhs error: %w", leftErr)
	}
	rightVal, rightErr := n.Right.Eval()
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
