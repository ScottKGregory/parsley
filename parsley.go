// Package parsley a parser/tokeniser used to evaluate expressions
package parsley

import (
	"fmt"

	"github.com/scottkgregory/parsley/internal/cache"
	"github.com/scottkgregory/parsley/internal/helpers"
	"github.com/scottkgregory/parsley/internal/nodes"
)

// Parser provides parsing and evaluation functionality
type Parser struct {
	cache cache.Store[string, nodes.Node]
}

// NewParser configures a new parser. If a cache it required one will be set up using github.com/dgraph-io/ristretto/v2
func NewParser(withCache bool) (m *Parser, err error) {
	m = &Parser{}
	if withCache {
		m.cache, err = cache.NewCache()
		if err != nil {
			return nil, err
		}
	} else {
		m.cache = cache.NewNoOpCache()
	}

	return m, nil
}

// Close releases the underlying resources
func (m *Parser) Close() {
	m.cache.Close()
}

// ParseAsBool is used to test whether the incoming data matches the given expression. Any numeric value over 0, or strings evaluating to true will match
func (m *Parser) ParseAsBool(str string, data map[string]any) (bool, error) {
	return parseAs(m, str, data, helpers.ToBool)
}

// ParseAsString will parse and evaluate the expression provided. Returning the result as a string, if the expression results in any other type it will be printed as a string
func (m *Parser) ParseAsString(str string, data map[string]any) (string, error) {
	return parseAs(m, str, data, helpers.ToString)
}

// ParseAsAny will parse and evaluate the expression provided. Returning the result as a whichever type is most appropriate
func (m *Parser) ParseAsAny(str string, data map[string]any) (any, error) {
	return parseAs(m, str, data, func(e any) (any, error) { return e, nil })
}

func parseAs[T any](m *Parser, str string, data map[string]any, converter func(e any) (T, error)) (T, error) {
	node, found := m.cache.Get(str)
	if !found {
		var err error
		node, err = parse(str)
		if err != nil {
			return *new(T), err
		}

		m.cache.Set(str, node)
	}

	val, err := node.Eval(data)
	if err != nil {
		return *new(T), fmt.Errorf("error evaluating expression: %w", err)
	}

	return converter(val)
}

// RegisterFunction registers a new function in the available set. Repeated calls will result in the latest one being registered
func RegisterFunction(name string, fun nodes.Function) {
	nodes.RegisterFunction(name, fun)
}
