package operations

import (
	"errors"
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		comp   string
		a, b   any
		result any
		err    error
	}{
		{comp: "+", a: 3, b: 3, result: 6, err: nil},
		{comp: "+", a: 3.4, b: 3, result: 6.4, err: nil},
		{comp: "+", a: 3.4, b: -3, result: 0.3999999999999999, err: nil},

		{comp: "/", a: 3, b: 3, result: 1, err: nil},
		{comp: "/", a: 3.4, b: 3, result: 1.1333333333333333, err: nil},
		{comp: "/", a: 3.4, b: -3, result: -1.1333333333333333, err: nil},

		{comp: "*", a: 3, b: 3, result: 9, err: nil},
		{comp: "*", a: 3.4, b: 3, result: 10.2, err: nil},
		{comp: "*", a: 3.4, b: -3, result: -10.2, err: nil},

		{comp: "-", a: 3, b: 3, result: 0, err: nil},
		{comp: "-", a: 3.4, b: 3, result: 0.3999999999999999, err: nil},
		{comp: "-", a: 3.4, b: -3, result: 6.4, err: nil},

		{comp: "^", a: 3, b: 3, result: 27, err: nil},
		{comp: "^", a: 3.4, b: 3, result: 39.303999999999995, err: nil},
		{comp: "^", a: 3.4, b: -3, result: 0.025442703032770204, err: nil},

		{comp: "&&", a: true, b: true, result: true, err: nil},
		{comp: "&&", a: true, b: false, result: false, err: nil},
		{comp: "&&", a: false, b: true, result: false, err: nil},
		{comp: "&&", a: false, b: false, result: false, err: nil},
		{comp: "&&", a: "blam", b: false, result: nil, err: errors.New("could not parse as boolean: string bool")},
		{comp: "&&", a: "blam", b: "blep", result: nil, err: errors.New("could not parse as boolean: string string")},
		{comp: "&&", a: true, b: "blep", result: nil, err: errors.New("could not parse as boolean: bool string")},
		{comp: "&&", a: "blam", b: 1, result: nil, err: errors.New("could not parse as boolean: string int")},
		{comp: "&&", a: "true", b: "true", result: true, err: nil},
		{comp: "&&", a: "true", b: "false", result: false, err: nil},
		{comp: "&&", a: "false", b: "true", result: false, err: nil},
		{comp: "&&", a: "false", b: "false", result: false, err: nil},

		{comp: "||", a: true, b: true, result: true, err: nil},
		{comp: "||", a: true, b: false, result: true, err: nil},
		{comp: "||", a: false, b: true, result: true, err: nil},
		{comp: "||", a: false, b: false, result: false, err: nil},
		{comp: "||", a: "blam", b: false, result: nil, err: errors.New("could not parse as boolean: string bool")},
		{comp: "||", a: "blam", b: "blep", result: nil, err: errors.New("could not parse as boolean: string string")},
		{comp: "||", a: true, b: "blep", result: nil, err: errors.New("could not parse as boolean: bool string")},
		{comp: "||", a: "blam", b: 1, result: nil, err: errors.New("could not parse as boolean: string int")},
		{comp: "||", a: "true", b: "true", result: true, err: nil},
		{comp: "||", a: "true", b: "false", result: true, err: nil},
		{comp: "||", a: "false", b: "true", result: true, err: nil},
		{comp: "||", a: "false", b: "false", result: false, err: nil},

		{comp: "<", a: "a", b: "b", result: true, err: nil},
		{comp: "<", a: "b", b: "a", result: false, err: nil},
		{comp: "<", a: 1, b: 2, result: true, err: nil},
		{comp: "<", a: 2, b: 1, result: false, err: nil},

		{comp: ">", a: "a", b: "b", result: false, err: nil},
		{comp: ">", a: "b", b: "a", result: true, err: nil},
		{comp: ">", a: 1, b: 2, result: false, err: nil},
		{comp: ">", a: 2, b: 1, result: true, err: nil},

		{comp: "==", a: "a", b: "b", result: false, err: nil},
		{comp: "==", a: "a", b: "a", result: true, err: nil},
		{comp: "==", a: "a", b: 2, result: nil, err: errors.New("only one side of comparison was a string: string int")},
		{comp: "==", a: 2, b: "a", result: nil, err: errors.New("only one side of comparison was a string: int string")},
		{comp: "==", a: 1, b: 2, result: false, err: nil},
		{comp: "==", a: 1, b: 1, result: true, err: nil},
		{comp: "=", a: "a", b: "b", result: false, err: nil},
		{comp: "=", a: "a", b: "a", result: true, err: nil},
		{comp: "=", a: "a", b: 2, result: nil, err: errors.New("only one side of comparison was a string: string int")},
		{comp: "=", a: 2, b: "a", result: nil, err: errors.New("only one side of comparison was a string: int string")},
		{comp: "=", a: 1, b: 2, result: false, err: nil},
		{comp: "=", a: 1, b: 1, result: true, err: nil},

		{comp: "£", a: 1, b: 1, result: nil, err: errors.New("unrecognised comparator: £")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v%s%v", tc.a, tc.comp, tc.b), func(t *testing.T) {
			op := &ComparisonOperation{tc.comp}
			actual, err := op.Calculate(tc.a, tc.b)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.result, actual)
		})
	}

}

func TestAddString(t *testing.T) {
	testCases := []struct {
		comp   string
		result string
	}{
		{comp: "+", result: "+"},
		{comp: "/", result: "/"},
		{comp: "*", result: "*"},
		{comp: "-", result: "-"},
		{comp: "^", result: "^"},
		{comp: "<", result: " < "},
		{comp: ">", result: " > "},
		{comp: "=", result: " = "},
		{comp: "==", result: " == "},
		{comp: "||", result: " || "},
		{comp: "&&", result: " && "},
	}
	for _, tc := range testCases {
		t.Run(tc.comp, func(t *testing.T) {
			actual := (&ComparisonOperation{tc.comp}).String()
			assert.Equal(t, tc.result, actual)
		})
	}
}
