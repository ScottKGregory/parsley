package nodes

import (
	"fmt"
	"strings"
)

// FunctionNode is the interface for an operation that executes a function
type FunctionNode struct {
	fun          func(args ...any) (any, error)
	FunctionName string
	Arguments    []Node
}

var _ Node = &FunctionNode{}

// NewFunctionNode creates a new function node
func NewFunctionNode(fun func(args ...any) (any, error), functionName string, arguments ...Node) *FunctionNode {
	return &FunctionNode{fun, functionName, arguments}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *FunctionNode) Eval(data map[string]any) (any, error) {
	// Evaluate all arguments
	argVals := make([]any, len(n.Arguments))
	for i, argument := range n.Arguments {
		var err error
		argVals[i], err = argument.Eval(data)
		if err != nil {
			return nil, fmt.Errorf("%w, error in argument %d: %w", ErrNodeEvalFailed, i, err)
		}
	}

	ret, err := n.fun(argVals...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrNodeEvalFailed, err)
	}

	return ret, nil
}

// String returns the string representation
func (n *FunctionNode) String() string {
	args := []string{}
	for _, n := range n.Arguments {
		args = append(args, n.String())
	}

	return fmt.Sprintf("%s(%s)", n.FunctionName, strings.Join(args, ", "))
}
