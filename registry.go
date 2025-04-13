package parsley

import (
	"fmt"
	"math"

	"github.com/scottkgregory/parsley/internal/helpers"
	"github.com/scottkgregory/parsley/internal/nodes"
)

// Function defines the shape of a function that can be called inside an expression
type Function func(args ...any) (any, error)
type UnaryNodeFunc func(right nodes.Node) nodes.Node
type BinaryNodeFunc func(left, right nodes.Node) nodes.Node
type registry struct {
	knownTokens []string
	unaryNodes  map[string]UnaryNodeFunc
	binaryNodes map[string]BinaryNodeFunc
	functions   map[string]Function
}

func newRegistry() *registry {
	return &registry{
		[]string{`+`, `-`, `*`, `^`, `/`, `(`, `)`, `""`, `,`, `==`, `>`, `<`, `&&`, `||`},
		map[string]UnaryNodeFunc{},
		map[string]BinaryNodeFunc{},
		map[string]Function{
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
					match, err := nodes.Calculate("==", v.(map[string]any)[args[1].(string)], args[2])
					if match.(bool) || err != nil {
						return match, err
					}
				}

				return false, nil
			},
		},
	}
}

func (p *Parser) RegisterUnaryNode(token string, fun UnaryNodeFunc) {
	p.Registry.knownTokens = append(p.Registry.knownTokens, token)
	p.Registry.unaryNodes[token] = fun
}

func (p *Parser) RegisterBinaryNode(token string, fun BinaryNodeFunc) {
	p.Registry.knownTokens = append(p.Registry.knownTokens, token)
	p.Registry.binaryNodes[token] = fun
}

// RegisterFunction registers a new function in the available set. Repeated calls will result in the latest one being registered
func (p *Parser) RegisterFunction(name string, fun Function) {
	p.Registry.functions[name] = fun
}
