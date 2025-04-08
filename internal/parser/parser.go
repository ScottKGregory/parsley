// Package parser is a tokenizer/parser to evaluate matcher config
package parser

import (
	"errors"
	"fmt"

	"github.com/scottkgregory/parsley/internal/parser/nodes"
	"github.com/scottkgregory/parsley/internal/parser/operations"
)

type parser struct {
	tokenizer *tokenizer
	data      map[string]any
}

func Parse(str string, data map[string]any) (nodes.Node, error) {
	t, err := newTokenizer(str)
	if err != nil {
		return nil, err
	}

	return parseTokens(t, data)
}

func parseTokens(tokenizer *tokenizer, data map[string]any) (nodes.Node, error) {
	parser := parser{tokenizer, data}
	return parser.parseExpression()
}

func (p *parser) parseExpression() (nodes.Node, error) {
	expr, err := p.parseAddSubtract()
	if err != nil {
		return nil, err
	}

	// Check everything was consumed
	if p.tokenizer.Token != eof {
		return nil, errors.New("unexpected characters at end of expression")
	}

	return expr, nil
}

func (p *parser) parseAddSubtract() (nodes.Node, error) {
	// Parse the left hand side
	left, err := p.parseMultiplyDivide()
	if err != nil {
		return nil, err
	}

	for {
		// Work out the operator
		var op nodes.BinaryNodeOp
		switch p.tokenizer.Token {
		case add:
			op = &operations.AddOperation{}
		case subtract:
			op = &operations.SubtractOperation{}
		}

		// Binary operator found?
		if op == nil {
			return left, nil
		}

		// Skip the operator
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		// Parse the right hand side of the expression
		right, err := p.parseMultiplyDivide()
		if err != nil {
			return nil, err
		}

		// Create a binary node and use it as the left-hand side from now on
		left = nodes.NewBinaryNode(left, right, op)
	}
}

func (p *parser) parseMultiplyDivide() (nodes.Node, error) {
	// Parse the left hand side
	left, err := p.parseUnary()
	if err != nil {
		return nil, err
	}

	for {
		// Work out the operator
		var op nodes.BinaryNodeOp
		switch p.tokenizer.Token {
		case multiply:
			op = &operations.MultiplyOperation{}
		case divide:
			op = &operations.DivideOperation{}
		case power:
			op = &operations.PowerOperation{}
		case equal:
			op = &operations.ComparisonOperation{Comparator: "=="}
		case greaterThan:
			op = &operations.ComparisonOperation{Comparator: ">"}
		case lessThan:
			op = &operations.ComparisonOperation{Comparator: "<"}
		case and:
			op = &operations.ComparisonOperation{Comparator: "&&"}
		case or:
			op = &operations.ComparisonOperation{Comparator: "||"}
		}

		// Binary operator found?
		if op == nil {
			return left, nil
		}

		// Skip the operator
		err = p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		// Skip second equals
		if p.tokenizer.Token == equal || p.tokenizer.Token == and || p.tokenizer.Token == or {
			err = p.tokenizer.NextToken()
			if err != nil {
				return nil, err
			}
		}

		// Parse the right hand side of the expression
		var right, err = p.parseUnary()
		if err != nil {
			return nil, err
		}

		// Create a binary node and use it as the left-hand side from now on
		left = nodes.NewBinaryNode(left, right, op)
	}
}

func (p *parser) parseUnary() (nodes.Node, error) {
	// Positive operator is a no-op so just skip it
	if p.tokenizer.Token == add {
		// Skip
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}
		return p.parseUnary()
	}

	// Negative operator
	if p.tokenizer.Token == subtract {
		// Skip
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		// Parse right
		// Note p recurses to self to support negative of a negative
		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}

		// Create unary node
		return nodes.NewUnaryNode(right, &operations.NegateOperation{}), nil
	}

	// No positive/negative operator so parse a leaf node
	return p.parseLeaf()
}

func (p *parser) parseLeaf() (nodes.Node, error) {
	// Is it a number?
	if p.tokenizer.Token == number {
		node := nodes.NewNumberNode(p.tokenizer.Number)
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}
		return node, nil
	}

	// Parenthesis?
	if p.tokenizer.Token == openParens {
		// Skip '('
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		// Parse a top-level expression
		node, err := p.parseAddSubtract()
		if err != nil {
			return nil, err
		}

		// Check and skip ')'
		if p.tokenizer.Token != closeParens {
			return nil, errors.New("missing close parenthesis")
		}

		err = p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		// Return
		return node, nil
	}

	// Quotes?
	// TODO: This does not allow for escaping quotes
	if p.tokenizer.Token == quote {
		// Skip '"'
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		s := ""
		for {
			err := p.tokenizer.NextToken()
			if err != nil {
				return nil, err
			}

			s += p.tokenizer.Identifier

			// Check and skip '"'
			if p.tokenizer.Token == quote {
				err := p.tokenizer.NextToken()
				if err != nil {
					return nil, err
				}
				return nodes.NewStringNode(s), nil
			}
		}
	}

	// Variable
	if p.tokenizer.Token == identifier {
		// Capture the name and skip it
		name := p.tokenizer.Identifier
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		// Parens indicate a function call, otherwise just a variable
		if p.tokenizer.Token == openParens {
			// Function call

			// Skip parens
			err := p.tokenizer.NextToken()
			if err != nil {
				return nil, err
			}

			// Parse arguments
			var arguments = []nodes.Node{}
			for {
				// Parse argument and add to list
				n, err := p.parseAddSubtract()
				if err != nil {
					return nil, err
				}

				arguments = append(arguments, n)

				// Is there another argument?
				if p.tokenizer.Token == comma {
					err = p.tokenizer.NextToken()
					if err != nil {
						return nil, err
					}
					continue
				}

				// Get out
				break
			}

			// Check and skip ')'
			if p.tokenizer.Token != closeParens {
				return nil, errors.New("missing close parenthesis")
			}

			err = p.tokenizer.NextToken()
			if err != nil {
				return nil, err
			}

			// Create the function call node
			return nodes.NewFunctionNode(name, arguments, p.data), nil
		}

		return nodes.NewVariableNode(name, p.data), nil
	}

	// Don't Understand
	return nil, fmt.Errorf("unexpected token: %s", p.tokenizer.Token)
}
