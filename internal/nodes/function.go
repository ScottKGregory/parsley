package nodes

import (
	"fmt"
	"math"
	"strings"

	"github.com/scottkgregory/parsley/internal/helpers"
)

type Function func(args ...any) (any, error)

var functions map[string]Function = map[string]Function{
	"ceil": func(args ...any) (any, error) {
		x, err := helpers.ToFloat64(args[0])
		if err != nil {
			return nil, fmt.Errorf("error calling function ceil: %w", err)
		}
		return math.Ceil(x), nil
	},
	"floor": func(args ...any) (any, error) {
		x, err := helpers.ToFloat64(args[0])
		if err != nil {
			return nil, fmt.Errorf("error calling function floor: %w", err)
		}
		return math.Floor(x), nil
	},
	"round": func(args ...any) (any, error) {
		x, err := helpers.ToFloat64(args[0])
		if err != nil {
			return nil, fmt.Errorf("error calling function round: %w", err)
		}
		return math.Round(x), nil
	},
	"truncate": func(args ...any) (any, error) {
		x, err := helpers.ToFloat64(args[0])
		if err != nil {
			return nil, fmt.Errorf("error calling function truncate: %w", err)
		}
		return math.Trunc(x), nil
	},
	"absolute": func(args ...any) (any, error) {
		x, err := helpers.ToFloat64(args[0])
		if err != nil {
			return nil, fmt.Errorf("error calling function absolute: %w", err)
		}
		return math.Abs(x), nil
	},
	"contains_any": func(args ...any) (any, error) {
		arr := args[0].([]any)
		for _, v := range arr {
			match, err := Calculate("==", v.(map[string]any)[args[1].(string)], args[2])
			if match.(bool) || err != nil {
				return match, err
			}
		}

		return false, nil
	},
}

// RegisterFunction registers a new function in the available set. Repeated calls will result in the latest one being registered
func RegisterFunction(name string, fun Function) {
	functions[name] = fun
}

// FunctionNode is the interface for an operation that executes a function
type FunctionNode struct {
	FunctionName string
	Arguments    []Node
}

var _ Node = &FunctionNode{}

// NewFunctionNode creates a new function node
func NewFunctionNode(functionName string, arguments ...Node) *FunctionNode {
	return &FunctionNode{functionName, arguments}
}

// Eval runs the appropriate logic to evaluate the node and produce a single result
func (n *FunctionNode) Eval(data map[string]any) (any, error) {
	// Evaluate all arguments
	argVals := make([]any, len(n.Arguments))
	for i, argument := range n.Arguments {
		var err error
		argVals[i], err = argument.Eval(data)
		if err != nil {
			return nil, fmt.Errorf("error in argument %d: %w", i, err)
		}
	}

	f, ok := functions[n.FunctionName]
	if !ok {
		return nil, fmt.Errorf("function %s not found", n.FunctionName)
	}

	return f(argVals...)
}

// String returns the string representation
func (n *FunctionNode) String() string {
	args := []string{}
	for _, n := range n.Arguments {
		args = append(args, n.String())
	}

	return fmt.Sprintf("%s(%s)", n.FunctionName, strings.Join(args, ", "))
}
