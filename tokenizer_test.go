package parsley

import (
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input  string
		tokens []string
	}{
		{
			input:  "1+2",
			tokens: []string{number, "+", number, eof},
		},
		{
			input:  "10-1",
			tokens: []string{number, "-", number, eof},
		},
		{
			input:  "3/3",
			tokens: []string{number, "/", number, eof},
		},
		{
			input:  "(3+3)/3",
			tokens: []string{"(", number, "+", number, ")", "/", number, eof},
		},
		{
			input:  "(3/3)+5",
			tokens: []string{"(", number, "/", number, ")", "+", number, eof},
		},
		{
			input:  "5 ==         5",
			tokens: []string{number, "==", number, eof},
		},
		{
			input:  "foo == 5",
			tokens: []string{identifier, "==", number, eof},
		},
		{
			input:  `foo == "hello"`,
			tokens: []string{identifier, "==", `"`, identifier, `"`, eof},
		},
		{
			input:  "foo.bar.baz == \"hello\"",
			tokens: []string{identifier, "==", `"`, identifier, `"`, eof},
		},
		{
			input:  "foo.bar.baz == 6",
			tokens: []string{identifier, "==", number, eof},
		},
		{
			input:  `ceil(2.1)`,
			tokens: []string{identifier, "(", number, ")", eof},
		},
		{
			input:  `1 > 2`,
			tokens: []string{number, ">", number, eof},
		},
		{
			input:  `2 < 1`,
			tokens: []string{number, "<", number, eof},
		},
		{
			input:  `foo < 2`,
			tokens: []string{identifier, "<", number, eof},
		},
		{
			input:  `(foo < 2) && (foo > 1)`,
			tokens: []string{"(", identifier, "<", number, ")", "&&", "(", identifier, ">", number, ")", eof},
		},
		{
			input:  `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			tokens: []string{"(", identifier, "==", `"`, identifier, `"`, ")", "&&", "(", identifier, "==", `"`, identifier, `"`, ")", eof},
		},
		{
			input:  `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			tokens: []string{"(", identifier, "==", `"`, identifier, `"`, ")", "&&", identifier, "(", identifier, ",", `"`, identifier, `"`, ",", `"`, identifier, `"`, ")", eof},
		},
		{
			input:  `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			tokens: []string{"(", identifier, "==", `"`, identifier, `"`, ")", "&&", identifier, "(", identifier, ",", `"`, identifier, `"`, ",", number, ")", eof},
		},
		{
			input:  `(object_attributes.state == "opened") || contains_any(object_attributes.labels, "title", 70)`,
			tokens: []string{"(", identifier, "==", `"`, identifier, `"`, ")", "||", identifier, "(", identifier, ",", `"`, identifier, `"`, ",", number, ")", eof},
		},
		{
			input:  `object_attributes.state`,
			tokens: []string{identifier, eof},
		},
		{
			input:  `-2`,
			tokens: []string{"-", number, eof},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(tt *testing.T) {
			tok, err := newTokenizer(tc.input, newRegistry())
			assert.Equal(t, nil, err)

			toks := []string{}
			for tok.Token != eof {
				toks = append(toks, tok.Token)
				tok.NextToken()
			}
			toks = append(toks, tok.Token)

			assert.Equal(t, tc.tokens, toks)
		})
	}
}
