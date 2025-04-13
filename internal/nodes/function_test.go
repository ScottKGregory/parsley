package nodes

import (
	"errors"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestFunctionNode(t *testing.T) {
	data1 := map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}}

	testCases := []struct {
		name         string
		data         map[string]any
		args         []Node
		err          error
		result       any
		stringResult string
	}{
		{"ceil", nil, []Node{NewMockNode(nil, 2.1, nil, "2.1")}, nil, float64(3), "ceil(2.1)"},
		{"floor", nil, []Node{NewMockNode(nil, 2.1, nil, "2.1")}, nil, float64(2), "floor(2.1)"},
		{"round", nil, []Node{NewMockNode(nil, 2.1, nil, "2.1")}, nil, float64(2), "round(2.1)"},
		{"truncate", nil, []Node{NewMockNode(nil, 2.1, nil, "2.1")}, nil, float64(2), "truncate(2.1)"},
		{"absolute", nil, []Node{NewMockNode(nil, -2.1, nil, "2.1")}, nil, 2.1, "absolute(2.1)"},

		{"ceil", nil, []Node{NewMockNode(nil, "2.1", nil, "2.1")}, nil, float64(3), "ceil(2.1)"},
		{"floor", nil, []Node{NewMockNode(nil, "2.1", nil, "2.1")}, nil, float64(2), "floor(2.1)"},
		{"round", nil, []Node{NewMockNode(nil, "2.1", nil, "2.1")}, nil, float64(2), "round(2.1)"},
		{"truncate", nil, []Node{NewMockNode(nil, "2.1", nil, "2.1")}, nil, float64(2), "truncate(2.1)"},
		{"absolute", nil, []Node{NewMockNode(nil, "-2.1", nil, "2.1")}, nil, 2.1, "absolute(2.1)"},

		{
			"ceil",
			nil,
			[]Node{NewMockNode(nil, "bleh", nil, "2.1")},
			errors.New(`node evaluation failed: error calling function ceil: error parsing value as float, could not parse string 'bleh'`),
			nil,
			"ceil(2.1)",
		},
		{
			"floor",
			nil,
			[]Node{NewMockNode(nil, "bleh", nil, "2.1")},
			errors.New(`node evaluation failed: error calling function floor: error parsing value as float, could not parse string 'bleh'`),
			nil,
			"floor(2.1)",
		},
		{
			"round",
			nil,
			[]Node{NewMockNode(nil, "bleh", nil, "2.1")},
			errors.New(`node evaluation failed: error calling function round: error parsing value as float, could not parse string 'bleh'`),
			nil,
			"round(2.1)",
		},
		{
			"truncate",
			nil,
			[]Node{NewMockNode(nil, "bleh", nil, "2.1")},
			errors.New(`node evaluation failed: error calling function truncate: error parsing value as float, could not parse string 'bleh'`),
			nil,
			"truncate(2.1)",
		},
		{
			"absolute",
			nil,
			[]Node{NewMockNode(nil, "bleh", nil, "2.1")},
			errors.New(`node evaluation failed: error calling function absolute: error parsing value as float, could not parse string 'bleh'`),
			nil,
			"absolute(2.1)",
		},
		{
			"absolute",
			nil,
			[]Node{NewMockNode(nil, 2.1, errors.New("node evaluation failed: uh oh"), "2.1")},
			errors.New(`node evaluation failed, error in argument 0: node evaluation failed: uh oh`),
			nil,
			"absolute(2.1)",
		},

		{
			"contains_any",
			data1,
			[]Node{
				NewMockNode(data1, []any{map[string]any{"title": "automated"}}, nil, "object_attributes.labels"),
				NewMockNode(data1, "title", nil, `"title"`),
				NewMockNode(data1, "automated", nil, `"automated"`),
			},
			nil,
			true,
			`contains_any(object_attributes.labels, "title", "automated")`,
		},
		{
			"contains_any",
			data1,
			[]Node{
				NewMockNode(data1, []any{map[string]any{"title": "automated"}}, nil, "object_attributes.labels"),
				NewMockNode(data1, "title", nil, `"title"`),
				NewMockNode(data1, "manual", nil, `"manual"`),
			},
			nil,
			false,
			`contains_any(object_attributes.labels, "title", "manual")`,
		},

		{
			"not_found",
			data1,
			[]Node{},
			errors.New("node evaluation failed: function not found 'not_found'"),
			nil,
			`not_found()`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n := NewFunctionNode(tc.name, tc.args...)

			res, err := n.Eval(tc.data)
			assert.ErrorEqual(t, tc.err, err)
			assert.Equal(t, tc.result, res)

			assert.Equal(t, tc.stringResult, n.String())
		})
	}
}
