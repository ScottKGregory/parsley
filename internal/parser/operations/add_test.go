package operations

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name   string
		a, b   any
		result any
	}{
		{a: 3, b: 3, result: 6},
		{a: 3.4, b: 3, result: 6.4},
		{a: 3.4, b: -3, result: 0.3999999999999999},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v+%v", tc.a, tc.b), func(t *testing.T) {
			add := &AddOperation{}
			actual, err := add.Calculate(tc.a, tc.b)
			assertNil(t, err)
			assertEqual(t, tc.result, actual)
		})
	}

}

func TestAddString(t *testing.T) {
	assertEqual(t, "+", (&AddOperation{}).String())
}

func assertNil(t *testing.T, val any) {
	if val != nil {
		fmt.Printf("expected nil but got %v\n", val)
		t.Fail()
	}
}

func assertEqual(t *testing.T, expected, actual any) {
	if fmt.Sprint(expected) != fmt.Sprint(actual) {
		fmt.Printf("expected '%v' but got '%v'\n", expected, actual)
		t.Fail()
	}
}
