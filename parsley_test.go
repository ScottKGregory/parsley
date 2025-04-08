package parsley

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/parser"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		data           map[string]any
		expected       any
		expectedString string
		expectedMatch  bool
	}{
		{
			name:           "basic addition",
			input:          "1+1",
			data:           map[string]any{},
			expected:       2,
			expectedString: `1.00+1.00`,
			expectedMatch:  true,
		},
		{
			name:           "basic subtraction",
			input:          "10-1",
			data:           map[string]any{},
			expected:       9,
			expectedString: `10.00-1.00`,
			expectedMatch:  true,
		},
		{
			name:           "basic division",
			input:          "3/3",
			data:           map[string]any{},
			expected:       1,
			expectedString: `3.00/3.00`,
			expectedMatch:  true,
		},
		{
			name:           "basic division with brackets",
			input:          "(3+3)/3",
			data:           map[string]any{},
			expected:       2,
			expectedString: `3.00+3.00/3.00`,
			expectedMatch:  true,
		},
		{
			name:           "basic division with brackets",
			input:          "(3/3)+5",
			data:           map[string]any{},
			expected:       6,
			expectedString: `3.00/3.00+5.00`,
			expectedMatch:  true,
		},
		{
			name:           "const equality int",
			input:          "5 == 5",
			data:           map[string]any{},
			expected:       true,
			expectedString: `5.00 == 5.00`,
			expectedMatch:  true,
		},
		{
			name:           "variable equality int",
			input:          "foo == 5",
			data:           map[string]any{"foo": 5},
			expected:       true,
			expectedString: `foo == 5.00`,
			expectedMatch:  true,
		},
		{
			name:           "variable equality int (not equal)",
			input:          "foo == 5",
			data:           map[string]any{"foo": 6},
			expected:       false,
			expectedString: `foo == 5.00`,
			expectedMatch:  false,
		},
		{
			name:           "variable equality string",
			input:          `foo == "hello"`,
			data:           map[string]any{"foo": "hello"},
			expected:       true,
			expectedString: `foo == "hello"`,
			expectedMatch:  true,
		},
		{
			name:           "deep variable equality string",
			input:          "foo.bar.baz == \"hello\"",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "hello"}}},
			expected:       true,
			expectedString: `foo.bar.baz == "hello"`,
			expectedMatch:  true,
		},
		{
			name:           "deep variable equality int",
			input:          "foo.bar.baz == 6",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": 6}}},
			expected:       true,
			expectedString: `foo.bar.baz == 6.00`,
			expectedMatch:  true,
		},
		{
			name:           "math ceil",
			input:          `ceil(2.10)`,
			data:           map[string]any{},
			expected:       3,
			expectedString: `ceil(2.10)`,
			expectedMatch:  true,
		},
		{
			name:           "math floor",
			input:          `floor(2.90)`,
			data:           map[string]any{},
			expected:       2,
			expectedString: `floor(2.90)`,
			expectedMatch:  true,
		},
		{
			name:           "math round up",
			input:          `round(2.90)`,
			data:           map[string]any{},
			expected:       3,
			expectedString: `round(2.90)`,
			expectedMatch:  true,
		},
		{
			name:           "math round down",
			input:          `round(2.49)`,
			data:           map[string]any{},
			expected:       2,
			expectedString: `round(2.49)`,
			expectedMatch:  true,
		},
		{
			name:           "math truncate",
			input:          `truncate(2.90)`,
			data:           map[string]any{},
			expected:       2,
			expectedString: `truncate(2.90)`,
			expectedMatch:  true,
		},
		{
			name:           "math absolute",
			input:          `absolute(2.90)`,
			data:           map[string]any{},
			expected:       2.9,
			expectedString: `absolute(2.90)`,
			expectedMatch:  true,
		},
		{
			name:           "greater than",
			input:          `1 > 2`,
			data:           map[string]any{},
			expected:       false,
			expectedString: `1.00 > 2.00`,
			expectedMatch:  false,
		},
		{
			name:           "greater than",
			input:          `2 > 1`,
			data:           map[string]any{},
			expected:       true,
			expectedString: `2.00 > 1.00`,
			expectedMatch:  true,
		},
		{
			name:           "less than",
			input:          `2.00 < 1.00`,
			data:           map[string]any{},
			expected:       false,
			expectedString: `2.00 < 1.00`,
			expectedMatch:  false,
		},
		{
			name:           "less than",
			input:          `1.00 < 2.00`,
			data:           map[string]any{},
			expected:       true,
			expectedString: `1.00 < 2.00`,
			expectedMatch:  true,
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 1.5},
			expected:       true,
			expectedString: `foo < 2.00 && foo > 1.00`,
			expectedMatch:  true,
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 4},
			expected:       false,
			expectedString: `foo < 2.00 && foo > 1.00`,
			expectedMatch:  false,
		},
		{
			name:           "gitlab sample match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "automated"}}},
			expected:       true,
			expectedString: `object_attributes.state == "opened" && object_attributes.labels.title == "automated"`,
			expectedMatch:  true,
		},
		{
			name:           "gitlab sample no match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "made by gary"}}},
			expected:       false,
			expectedString: `object_attributes.state == "opened" && object_attributes.labels.title == "automated"`,
			expectedMatch:  false,
		},
		{
			name:           "contains_any (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}},
			expected:       true,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", "automated")`,
			expectedMatch:  true,
		},
		{
			name:           "contains_any (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "made by gary"}}}},
			expected:       false,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", "automated")`,
			expectedMatch:  false,
		},
		{
			name:           "contains_any number (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 70}}}},
			expected:       true,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", 70.00)`,
			expectedMatch:  true,
		},
		{
			name:           "contains_any number (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expected:       false,
			expectedString: `object_attributes.state == "opened" && contains_any(object_attributes.labels, "title", 70.00)`,
			expectedMatch:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := parser.Parse(tc.input, tc.data)
			assertNil(tt, err)
			if err != nil {
				return
			}

			val, err := actual.Eval()
			assertNil(t, err)
			assertEqual(tt, fmt.Sprint(tc.expected), fmt.Sprint(val))
			assertEqual(tt, tc.expectedString, actual.String())
		})
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("matcher_%s", tc.name), func(tt *testing.T) {
			matcher, err := NewMatcher(true)
			assertNil(t, err)

			actual, err := matcher.Match(tc.input, tc.data)
			assertEqual(tt, tc.expectedMatch, actual)
			assertNil(t, err)
		})
	}
}

func assertNil(t *testing.T, val any) {
	if val != nil {
		fmt.Printf("expected nil but got %v\n", val)
		t.Fail()
	}
}

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		fmt.Printf("expected '%v' but got '%v'\n", expected, actual)
		t.Fail()
	}
}
