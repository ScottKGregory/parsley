package nodes

import (
	"fmt"
)

// StringNode is a node used to store a string
type StringNode struct {
	StringValue string
}

var _ Node = &StringNode{}

// NewStringNode creates a new string node
func NewStringNode(stringValue string) *StringNode {
	return &StringNode{stringValue}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *StringNode) Eval(data map[string]any) (any, error) {
	return n.StringValue, nil

}

// String returns the string representation
func (n *StringNode) String() string {
	return fmt.Sprintf("\"%s\"", n.StringValue)
}
