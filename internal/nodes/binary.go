package nodes

import (
	"fmt"
	"math"
	"strings"

	"github.com/scottkgregory/parsley/internal/helpers"
)

// BinaryNode is a node that has both a left and right side
type BinaryNode struct {
	Left  Node
	Right Node
	op    string
}

var _ Node = &BinaryNode{}

// NewBinaryNode creates a new binary node
func NewBinaryNode(left, right Node, op string) *BinaryNode {
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
	return Calculate(n.op, leftVal, rightVal)
}

// String returns the string representation
func (n *BinaryNode) String() string {
	f := "%s%s%s"
	if n.op == "<" || n.op == ">" || n.op == "=" || n.op == "==" || n.op == "||" || n.op == "&&" {
		f = "%s %s %s"
	}

	return fmt.Sprintf(f, n.Left.String(), n.op, n.Right.String())
}

func Calculate(op string, a, b any) (any, error) {
	if op == "||" || op == "&&" {
		x, err := helpers.ToBool(a)
		if err != nil {
			return nil, fmt.Errorf("could not parse as boolean: %T %T", a, b)
		}

		y, err := helpers.ToBool(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse as boolean: %T %T", a, b)
		}

		if op == "||" {
			return x || y, nil
		}

		return x && y, nil
	}

	x, aOk := a.(string)
	y, bOk := b.(string)
	if aOk && bOk {
		switch op {
		case "<":
			return strings.Compare(x, y) < 0, nil
		case ">":
			return strings.Compare(x, y) > 0, nil
		case "=", "==":
			return x == y, nil
		}
	}

	if aOk || bOk {
		return nil, fmt.Errorf("only one side of comparison was a string: %T %T", a, b)
	}

	aa, aErr := helpers.ToFloat64(a)
	if aErr != nil {
		return nil, fmt.Errorf("error in lhs: %w", aErr)
	}

	bb, bErr := helpers.ToFloat64(b)
	if bErr != nil {
		return nil, fmt.Errorf("error in rhs: %w", bErr)
	}

	switch op {
	case "<":
		return aa < bb, nil
	case ">":
		return aa > bb, nil
	case "=", "==":
		return aa == bb, nil
	case "+":
		return aa + bb, nil
	case "/":
		return aa / bb, nil
	case "*":
		return aa * bb, nil
	case "-":
		return aa - bb, nil
	case "^":
		return math.Pow(aa, bb), nil
	}

	return nil, fmt.Errorf("unrecognised op: %s", string(op))
}
