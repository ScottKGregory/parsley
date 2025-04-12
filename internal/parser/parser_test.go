package parser

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		data           map[string]any
		expected       any
		expectedString string
	}{
		{
			name:           "basic addition",
			input:          "1+1",
			data:           map[string]any{},
			expected:       2,
			expectedString: `1+1`,
		},
		{
			name:           "basic subtraction",
			input:          "10-1",
			data:           map[string]any{},
			expected:       9,
			expectedString: `10-1`,
		},
		{
			name:           "basic division",
			input:          "3/3",
			data:           map[string]any{},
			expected:       1,
			expectedString: `3/3`,
		},
		{
			name:           "basic division with brackets",
			input:          "(3+3)/3",
			data:           map[string]any{},
			expected:       2,
			expectedString: `3+3/3`,
		},
		{
			name:           "basic division with brackets",
			input:          "(3/3)+5",
			data:           map[string]any{},
			expected:       6,
			expectedString: `3/3+5`,
		},
		{
			name:           "const equality int",
			input:          "5 == 5",
			data:           map[string]any{},
			expected:       true,
			expectedString: `5 == 5`,
		},
		{
			name:           "variable equality int",
			input:          "foo == 5",
			data:           map[string]any{"foo": 5},
			expected:       true,
			expectedString: `foo == 5`,
		},
		{
			name:           "variable equality int (not equal)",
			input:          "foo == 5",
			data:           map[string]any{"foo": 6},
			expected:       false,
			expectedString: `foo == 5`,
		},
		{
			name:           "variable equality string",
			input:          `foo == "hello"`,
			data:           map[string]any{"foo": "hello"},
			expected:       true,
			expectedString: `foo == "hello"`,
		},
		{
			name:           "deep variable equality string",
			input:          "foo.bar.baz == \"hello\"",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "hello"}}},
			expected:       true,
			expectedString: `foo.bar.baz == "hello"`,
		},
		{
			name:           "deep variable equality int",
			input:          "foo.bar.baz == 6",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": 6}}},
			expected:       true,
			expectedString: `foo.bar.baz == 6`,
		},
		{
			name:           "math ceil",
			input:          `ceil(2.1)`,
			data:           map[string]any{},
			expected:       3,
			expectedString: `ceil(2.1)`,
		},
		{
			name:           "math floor",
			input:          `floor(2.9)`,
			data:           map[string]any{},
			expected:       2,
			expectedString: `floor(2.9)`,
		},
		{
			name:           "math round up",
			input:          `round(2.9)`,
			data:           map[string]any{},
			expected:       3,
			expectedString: `round(2.9)`,
		},
		{
			name:           "math round down",
			input:          `round(2.49)`,
			data:           map[string]any{},
			expected:       2,
			expectedString: `round(2.49)`,
		},
		{
			name:           "math truncate",
			input:          `truncate(2.9)`,
			data:           map[string]any{},
			expected:       2,
			expectedString: `truncate(2.9)`,
		},
		{
			name:           "math absolute",
			input:          `absolute(2.9)`,
			data:           map[string]any{},
			expected:       2.9,
			expectedString: `absolute(2.9)`,
		},
		{
			name:           "greater than",
			input:          `1 > 2`,
			data:           map[string]any{},
			expected:       false,
			expectedString: `1 > 2`,
		},
		{
			name:           "greater than",
			input:          `2 > 1`,
			data:           map[string]any{},
			expected:       true,
			expectedString: `2 > 1`,
		},
		{
			name:           "less than",
			input:          `2 < 1`,
			data:           map[string]any{},
			expected:       false,
			expectedString: `2 < 1`,
		},
		{
			name:           "less than",
			input:          `1 < 2`,
			data:           map[string]any{},
			expected:       true,
			expectedString: `1 < 2`,
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 1.5},
			expected:       true,
			expectedString: `foo < 2 && foo > 1`,
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 4},
			expected:       false,
			expectedString: `foo < 2 && foo > 1`,
		},
		{
			name:           "gitlab sample match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "automated"}}},
			expected:       true,
			expectedString: `object_attributes.state == "opened" && object_attributes.labels.title == "automated"`,
		},
		{
			name:           "gitlab sample no match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "made by gary"}}},
			expected:       false,
			expectedString: `object_attributes.state == "opened" && object_attributes.labels.title == "automated"`,
		},
		{
			name:           "contains_any (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}},
			expected:       true,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", "automated")`,
		},
		{
			name:           "contains_any (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "made by gary"}}}},
			expected:       false,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", "automated")`,
		},
		{
			name:           "contains_any number (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 70}}}},
			expected:       true,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", 70)`,
		},
		{
			name:           "contains_any number (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expected:       false,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", 70)`,
		},
		{
			name:           "order of operations",
			input:          `36/6*3+2^2-(3+5)`,
			data:           nil,
			expected:       14,
			expectedString: "36/6*3+2^2-3+5",
		},
		{
			name:           "order of operations",
			input:          `10 + (5 * 3 + 2)`,
			data:           nil,
			expected:       27,
			expectedString: "10+5*3+2",
		},
		{
			name:           "order of operations",
			input:          `15 + (30 / 2)`,
			data:           nil,
			expected:       30,
			expectedString: "15+30/2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := Parse(tc.input)
			assert.Nil(tt, err)
			if err != nil {
				return
			}

			val, err := actual.Eval(tc.data)
			assert.Nil(t, err)
			assert.Equal(tt, fmt.Sprint(tc.expected), fmt.Sprint(val))
			assert.Equal(tt, tc.expectedString, actual.String())
		})
	}
}
