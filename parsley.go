package parsley

import (
	"fmt"

	cache "github.com/scottkgregory/parsley/internal"
	"github.com/scottkgregory/parsley/internal/helpers"
	"github.com/scottkgregory/parsley/internal/parser"
	"github.com/scottkgregory/parsley/internal/parser/nodes"
)

type Matcher struct {
	cache cache.Store[string, nodes.Node]
}

func NewMatcher(withCache bool) (m *Matcher, err error) {
	m = &Matcher{}
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

func (m *Matcher) Close() {
	m.cache.Close()
}

// Match is used to test whether the incoming data matches the given expression. Any numeric value over 0, or strings evaluating to true will match
func (m *Matcher) Match(str string, data map[string]any) (bool, error) {
	node, found := m.cache.Get(str)
	if !found {
		var err error
		node, err = parser.Parse(str)
		if err != nil {
			return false, err
		}

		m.cache.Set(str, node)
	}

	val, err := node.Eval(data)
	if err != nil {
		return false, fmt.Errorf("error evaluating expression: %w", err)
	}

	return helpers.ToBool(val)
}
