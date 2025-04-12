package assert

import (
	"fmt"
	"testing"
)

func Nil(t *testing.T, val any) {
	if val != nil {
		fmt.Printf("expected nil but got %v\n", val)
		t.Fail()
	}
}

func Equal(t *testing.T, expected, actual any) {
	if fmt.Sprint(expected) != fmt.Sprint(actual) {
		fmt.Printf("expected '%v' but got '%v'\n", expected, actual)
		t.Fail()
	}
}
