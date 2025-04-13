// Package assert is used exclusively for asserting unit test output
package assert

import (
	"errors"
	"fmt"
	"testing"
)

// Nil asserts that the value provided is nil
func Nil(t *testing.T, val any) {
	if val != nil {
		t.Logf("expected nil but got %v\n", val)
		t.Fail()
	}
}

// Equal asserts that the two supplied values are equal
func Equal(t *testing.T, expected, actual any) {
	if fmt.Sprint(expected) != fmt.Sprint(actual) || fmt.Sprintf("%T", expected) != fmt.Sprintf("%T", actual) {
		t.Logf("expected '%v' (%T) but got '%v' (%T)\n", expected, expected, actual, actual)
		t.Fail()
	}
}

// ErrorEqual asserts that the two supplied errors are equal
func ErrorEqual(t *testing.T, expected, actual error) {
	if fmt.Sprint(expected) != fmt.Sprint(actual) {
		t.Logf("expected '%v' but got '%v'\n", expected, actual)
		t.Fail()
	}
}

// ErrorIs asserts that the actual error is the same type as the expected one
func ErrorIs(t *testing.T, expected, actual error) {
	if !errors.Is(actual, expected) {
		t.Logf("expected error to be '%s' but got '%s'\n", expected, actual)
		t.Fail()
	}
}
