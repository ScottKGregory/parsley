package parsley

import (
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		data          map[string]any
		expectedMatch bool
	}{
		{
			name:          "basic addition",
			input:         "1+1",
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "basic subtraction",
			input:         "10-1",
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "basic division",
			input:         "3/3",
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "basic division with brackets",
			input:         "(3+3)/3",
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "basic division with brackets",
			input:         "(3/3)+5",
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "const equality int",
			input:         "5 == 5",
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "variable equality int",
			input:         "foo == 5",
			data:          map[string]any{"foo": 5},
			expectedMatch: true,
		},
		{
			name:          "variable equality int (not equal)",
			input:         "foo == 5",
			data:          map[string]any{"foo": 6},
			expectedMatch: false,
		},
		{
			name:          "variable equality string",
			input:         `foo == "hello"`,
			data:          map[string]any{"foo": "hello"},
			expectedMatch: true,
		},
		{
			name:          "deep variable equality string",
			input:         "foo.bar.baz == \"hello\"",
			data:          map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "hello"}}},
			expectedMatch: true,
		},
		{
			name:          "deep variable equality int",
			input:         "foo.bar.baz == 6",
			data:          map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": 6}}},
			expectedMatch: true,
		},
		{
			name:          "math ceil",
			input:         `ceil(2.1)`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "math floor",
			input:         `floor(2.9)`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "math round up",
			input:         `round(2.9)`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "math round down",
			input:         `round(2.49)`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "math truncate",
			input:         `truncate(2.9)`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "math absolute",
			input:         `absolute(2.9)`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "greater than",
			input:         `1 > 2`,
			data:          map[string]any{},
			expectedMatch: false,
		},
		{
			name:          "greater than",
			input:         `2 > 1`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "less than",
			input:         `2 < 1`,
			data:          map[string]any{},
			expectedMatch: false,
		},
		{
			name:          "less than",
			input:         `1 < 2`,
			data:          map[string]any{},
			expectedMatch: true,
		},
		{
			name:          "less than and greater than",
			input:         `(foo < 2) && (foo > 1)`,
			data:          map[string]any{"foo": 1.5},
			expectedMatch: true,
		},
		{
			name:          "less than and greater than",
			input:         `(foo < 2) && (foo > 1)`,
			data:          map[string]any{"foo": 4},
			expectedMatch: false,
		},
		{
			name:          "gitlab sample match",
			input:         `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:          map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "automated"}}},
			expectedMatch: true,
		},
		{
			name:          "gitlab sample no match",
			input:         `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:          map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "made by gary"}}},
			expectedMatch: false,
		},
		{
			name:          "contains_any (true)",
			input:         `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:          map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}},
			expectedMatch: true,
		},
		{
			name:          "contains_any (false)",
			input:         `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:          map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "made by gary"}}}},
			expectedMatch: false,
		},
		{
			name:          "contains_any number (true)",
			input:         `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:          map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 70}}}},
			expectedMatch: true,
		},
		{
			name:          "contains_any number (false)",
			input:         `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:          map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expectedMatch: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("matcher_%s", tc.name), func(tt *testing.T) {
			matcher, err := NewMatcher(true)
			assert.Nil(t, err)

			actual, err := matcher.Match(tc.input, tc.data)
			assert.Equal(tt, tc.expectedMatch, actual)
			assert.Nil(t, err)
		})
	}
}
