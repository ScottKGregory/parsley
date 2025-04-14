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
	cache    cache.Store[string, nodes.Node]
	Registry *registry
}

// NewParser configures a new parser. If a cache it required one will be set up using github.com/dgraph-io/ristretto/v2
func NewParser(withCache bool) (m *Parser, err error) {
	m = &Parser{
		Registry: newRegistry(),
	}
	if withCache {
		m.cache, err = cache.NewCache()
		if err != nil {
			return nil, err //nolint:wrapcheck // Error is already wrapped
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

// ParseAsFloat will parse and evaluate the expression provided. Returning the result as a float64
func (m *Parser) ParseAsFloat(str string, data map[string]any) (float64, error) {
	return parseAs(m, str, data, helpers.ToFloat64)
}

// ParseAsAny will parse and evaluate the expression provided. Returning the result as a whichever type is most appropriate
func (m *Parser) ParseAsAny(str string, data map[string]any) (any, error) {
	return parseAs(m, str, data, func(e any) (any, error) { return e, nil })
}

func parseAs[T any](m *Parser, str string, data map[string]any, converter func(e any) (T, error)) (T, error) {
	node, found := m.cache.Get(str)
	if !found {
		var err error
		node, err = parse(str, m.Registry)
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

// ErrCacheSetup is returned when setting up the cache fails
const ErrCacheSetup = cache.ErrCacheSetup

// ErrComparisonFailed is returned when the comparison of two values fails
const ErrComparisonFailed = nodes.ErrComparisonFailed

// ErrNodeEvalFailed is returned when a node failed to evaluate correctly
const ErrNodeEvalFailed = nodes.ErrNodeEvalFailed

// TypesMatch check if the types of the two values are the same
func TypesMatch(a, b any) bool { return helpers.TypesMatch(a, b) }

// ToFloat64 attempts to convert the input value in to a float64. It will cast int/uint/flaot types, and attempt to parse strings as floats
func ToFloat64(input any) (float64, error) { return helpers.ToFloat64(input) }

// ToBool converts the input value in to a bool.
//
// - If the type is a number then any value over 0 will return true
// - Strings will be checked against known values
// - Strings not matching a known value will attempt to parse as a float
func ToBool(e any) (bool, error) { return helpers.ToBool(e) }
