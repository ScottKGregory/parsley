// Package assert is used exclusively for asserting unit test output
package assert

import (
	"fmt"
	"testing"
)

// Nil asserts that the value provided is nil
func Nil(t *testing.T, val any) {
	if val != nil {
		fmt.Printf("expected nil but got %v\n", val)
		t.Fail()
	}
}

// Equal asserts that the two supplied values are equal
func Equal(t *testing.T, expected, actual any) {
	if fmt.Sprint(expected) != fmt.Sprint(actual) || fmt.Sprintf("%T", expected) != fmt.Sprintf("%T", actual) {
		fmt.Printf("expected '%v' (%T) but got '%v' (%T)\n", expected, expected, actual, actual)
		t.Fail()
	}
}
