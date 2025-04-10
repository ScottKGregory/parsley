package nodes

import (
	"strings"
)

// VariableNode is a node used to store a variable reference
type VariableNode struct {
	VariableName string
}

var _ Node = &VariableNode{}

// NewVariableNode creates a new variable node
func NewVariableNode(variableName string) *VariableNode {
	return &VariableNode{variableName}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *VariableNode) Eval(data map[string]any) (any, error) {
	return getValue(strings.Split(n.VariableName, "."), data), nil
}

// String returns the string representation
func (n *VariableNode) String() string {
	return n.VariableName
}

func getValue(key []string, data map[string]any) any {
	if len(key) == 0 {
		return nil
	}

	x := data[key[0]]
	if x, ok := x.(map[string]any); ok {
		return getValue(key[1:], x)
	}

	return x
}
