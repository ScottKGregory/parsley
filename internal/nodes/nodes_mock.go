package nodes

import (
	"fmt"
	"reflect"
	"testing"
)

// MockNode is a mocked version of the Node interface, to be used in tests
type MockNode struct {
	evalCalled   bool
	stringCalled bool

	evalExpectedData map[string]any
	evalResult       any
	evalErr          error

	stringResult string
}

var _ Node = &MockNode{}

// NewMockNode creates a new instance of a MockNode
func NewMockNode(evalExpectedData map[string]any, evalResult any, evalErr error, stringResult string) *MockNode {
	return &MockNode{
		evalExpectedData: evalExpectedData,
		evalResult:       evalResult,
		evalErr:          evalErr,
		stringResult:     stringResult,
	}
}

// Eval mocks the Eval function
func (m *MockNode) Eval(data map[string]any) (any, error) {
	if !reflect.DeepEqual(m.evalExpectedData, data) {
		panic(fmt.Errorf("supplied data did not match in call to Eval expected: %#v, actual: %#v", m.evalExpectedData, data))
	}

	m.evalCalled = true
	return m.evalResult, m.evalErr
}

// String mocks the String function
func (m *MockNode) String() string {
	m.stringCalled = true
	return m.stringResult
}

// AssertEvalCalled checks that Eval has been called
func (m *MockNode) AssertEvalCalled(t *testing.T) {
	if !m.evalCalled {
		fmt.Println("expected call to Eval")
		t.Fail()
	}
}

// AssertStringCalled checks that String has been called
func (m *MockNode) AssertStringCalled(t *testing.T) {
	if !m.stringCalled {
		fmt.Println("expected call to String")
		t.Fail()
	}
}
