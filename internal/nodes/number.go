package nodes

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/helpers"
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
func (n *NumberNode) Eval(_ map[string]any) (any, error) {
	ret, err := helpers.ToFloat64(n.Number)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrNodeEvalFailed, err)
	}

	return ret, nil
}

// String returns the string representation
func (n *NumberNode) String() string {
	return fmt.Sprintf("%v", n.Number)
}
