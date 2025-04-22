package parsley

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
		expectedBool   *bool
		expectedAny    any
		expectedString any
	}{
		{
			name:           "basic negation",
			input:          "-2",
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    float64(-2),
			expectedString: "-2",
		},
		{
			name:           "negation",
			input:          "-(2*2)",
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    float64(-4),
			expectedString: "-4",
		},
		{
			name:           "basic addition",
			input:          "1+2",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(3),
			expectedString: "3",
		},
		{
			name:           "basic subtraction",
			input:          "10-1",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(9),
			expectedString: "9",
		},
		{
			name:           "basic division",
			input:          "3/3",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(1),
			expectedString: "1",
		},
		{
			name:           "basic division with brackets",
			input:          "(3+3)/3",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "basic division with brackets",
			input:          "(3/3)+5",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(6),
			expectedString: "6",
		},
		{
			name:           "const equality int",
			input:          "5 == 5",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "variable equality int",
			input:          "foo == 5",
			data:           map[string]any{"foo": 5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "variable equality int (not equal)",
			input:          "foo == 5",
			data:           map[string]any{"foo": 6},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "variable equality string",
			input:          `foo == "hello"`,
			data:           map[string]any{"foo": "hello"},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "deep variable equality string",
			input:          "foo.bar.baz == \"hello\"",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "hello"}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "deep variable equality int",
			input:          "foo.bar.baz == 6",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": 6}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "math ceil",
			input:          `ceil(2.1)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(3),
			expectedString: "3",
		},
		{
			name:           "math floor",
			input:          `floor(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "math round up",
			input:          `round(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(3),
			expectedString: "3",
		},
		{
			name:           "math round down",
			input:          `round(2.49)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "math truncate",
			input:          `truncate(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "math absolute",
			input:          `absolute(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    2.9,
			expectedString: "2.9",
		},
		{
			name:           "greater than",
			input:          `1 > 2`,
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "greater than",
			input:          `2 > 1`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than",
			input:          `2 < 1`,
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "less than",
			input:          `1 < 2`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than",
			input:          `foo < 2`,
			data:           map[string]any{"foo": 1.5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 1.5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 4},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "less than or greater than",
			input:          `(foo < 2) || (foo > 1)`,
			data:           map[string]any{"foo": 1.5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than or greater than",
			input:          `(foo < 2) || (foo > 1)`,
			data:           map[string]any{"foo": 4},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "gitlab sample match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "automated"}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "gitlab sample no match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "made by gary"}}},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "contains_any (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "contains_any (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "made by gary"}}}},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "contains_any number (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 70}}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "contains_any number (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "prints data",
			input:          `object_attributes.state`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expectedBool:   nil,
			expectedAny:    "opened",
			expectedString: "opened",
		},
		{
			name:           "equal, no data",
			input:          `event_type == "merge_request"`,
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
	}

	for _, tc := range testCases {
		if tc.expectedBool != nil {
			t.Run(fmt.Sprintf("as_bool_%s", tc.name), func(tt *testing.T) {
				matcher, err := NewParser(true)
				assert.Nil(t, err)

				actual, err := matcher.ParseAsBool(tc.input, tc.data)
				assert.Equal(tt, *tc.expectedBool, actual)
				assert.Nil(t, err)

				matcher.Close()
			})
		}

		t.Run(fmt.Sprintf("as_string_%s", tc.name), func(tt *testing.T) {
			matcher, err := NewParser(false)
			assert.Nil(t, err)

			actual, err := matcher.ParseAsString(tc.input, tc.data)
			assert.Equal(tt, tc.expectedString, actual)
			assert.Nil(t, err)

			matcher.Close()
		})

		t.Run(fmt.Sprintf("as_any_%s", tc.name), func(tt *testing.T) {
			matcher, err := NewParser(false)
			assert.Nil(t, err)

			actual, err := matcher.ParseAsAny(tc.input, tc.data)
			assert.Equal(tt, tc.expectedAny, actual)
			assert.Nil(t, err)

			matcher.Close()
		})
	}
}

func toPtr[T any](t T) *T {
	return &t
}
