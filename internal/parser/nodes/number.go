package nodes

import (
	"fmt"
)

// NumberNode is a node used to store a number
type NumberNode struct {
	Number any
}

var _ Node = &NumberNode{}

// NewNumberNode creates a new number node
func NewNumberNode(number any) *NumberNode {
	return &NumberNode{number}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *NumberNode) Eval(data map[string]any) (any, error) {
	return n.Number, nil
}

// String returns the string representation
func (n *NumberNode) String() string {
	return fmt.Sprintf("%.2f", n.Number)
}
