package parsley

import (
	"errors"
	"fmt"

	"github.com/scottkgregory/parsley/internal/nodes"
)

type parser struct {
	tokenizer *tokenizer
}

func parse(str string) (nodes.Node, error) {
	t, err := newTokenizer(str)
	if err != nil {
		return nil, err
	}

	return (&parser{t}).parseExpression()
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
		var op string
		switch p.tokenizer.Token {
		case "+", "-":
			op = p.tokenizer.Token
		}

		// Binary operator found?
		if op == "" {
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
		var op string
		switch p.tokenizer.Token {
		case "*", "/", "^", "==", "=", ">", "<", "&&", "&", "||", "|":
			op = p.tokenizer.Token
		}

		// Binary operator found?
		if op == "" {
			return left, nil
		}

		// Skip the operator
		err = p.tokenizer.NextToken()
		if err != nil {
			return nil, err
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
	if p.tokenizer.Token == "+" {
		// Skip
		err := p.tokenizer.NextToken()
		if err != nil {
			return nil, err
		}
		return p.parseUnary()
	}

	// Negative operator
	if p.tokenizer.Token == "-" {
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
		return nodes.NewUnaryNode(right, p.tokenizer.Token), nil
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
	if p.tokenizer.Token == "(" {
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
		if p.tokenizer.Token != ")" {
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
	if p.tokenizer.Token == `"` {
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
			if p.tokenizer.Token == `"` {
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
		if p.tokenizer.Token == "(" {
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
				if p.tokenizer.Token == "," {
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
			if p.tokenizer.Token != ")" {
				return nil, errors.New("missing close parenthesis")
			}

			err = p.tokenizer.NextToken()
			if err != nil {
				return nil, err
			}

			// Create the function call node
			return nodes.NewFunctionNode(name, arguments...), nil
		}

		return nodes.NewVariableNode(name), nil
	}

	// Don't Understand
	return nil, fmt.Errorf("unexpected token: %s", p.tokenizer.Token)
}
