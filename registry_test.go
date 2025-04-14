package parsley

import (
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
	"github.com/scottkgregory/parsley/internal/nodes"
)

func TestRegisterUnary(t *testing.T) {
	parser, err := NewParser(false)
	assert.Nil(t, err)
	parser.RegisterUnaryNode("£", func(_ nodes.Node) nodes.Node { return &testNode{} })

	actual, err := parser.ParseAsAny("£14", nil)
	assert.Equal(t, 12, actual)
	assert.Nil(t, err)

	parser.Close()
}

func TestRegisterBinary(t *testing.T) {
	parser, err := NewParser(false)
	assert.Nil(t, err)
	parser.RegisterBinaryNode("£", func(_, _ nodes.Node) nodes.Node { return &testNode{} })

	actual, err := parser.ParseAsAny("1£4", nil)
	assert.Equal(t, 12, actual)
	assert.Nil(t, err)

	parser.Close()
}

func TestRegiserFunction(t *testing.T) {
	testCases := []struct {
		name         string
		data         map[string]any
		args         []any
		err          error
		result       any
		stringResult string
	}{
		{"ceil", nil, []any{2.1}, nil, float64(3), "ceil(2.1)"},
		{"floor", nil, []any{2.1}, nil, float64(2), "floor(2.1)"},
		{"round", nil, []any{2.1}, nil, float64(2), "round(2.1)"},
		{"truncate", nil, []any{2.1}, nil, float64(2), "truncate(2.1)"},
		{"absolute", nil, []any{-2.1}, nil, 2.1, "absolute(2.1)"},

		{"ceil", nil, []any{"2.1"}, nil, float64(3), "ceil(2.1)"},
		{"floor", nil, []any{"2.1"}, nil, float64(2), "floor(2.1)"},
		{"round", nil, []any{"2.1"}, nil, float64(2), "round(2.1)"},
		{"truncate", nil, []any{"2.1"}, nil, float64(2), "truncate(2.1)"},
		{"absolute", nil, []any{"-2.1"}, nil, 2.1, "absolute(2.1)"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parser, err := NewParser(false)
			assert.Nil(t, err)

			actual, err := parser.Registry.functions[tc.name](tc.args...)
			assert.Equal(t, tc.result, actual)
			assert.ErrorEqual(t, tc.err, err)

			parser.Close()
		})
	}
}

type testNode struct{}

// Eval implements nodes.Node.
func (t *testNode) Eval(_ map[string]any) (any, error) {
	return 12, nil
}

// String implements nodes.Node.
func (t *testNode) String() string {
	return "£"
}
