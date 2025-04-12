package nodes

import (
	"fmt"
	"reflect"
	"testing"
)

type MockNode struct {
	evalCalled   bool
	stringCalled bool

	evalExpectedData map[string]any
	evalResult       any
	evalErr          error

	stringResult string
}

var _ Node = &MockNode{}

func NewMockNode(evalExpectedData map[string]any, evalResult any, evalErr error, stringResult string) *MockNode {
	return &MockNode{
		evalExpectedData: evalExpectedData,
		evalResult:       evalResult,
		evalErr:          evalErr,
		stringResult:     stringResult,
	}
}

func (m *MockNode) Eval(data map[string]any) (any, error) {
	if !reflect.DeepEqual(m.evalExpectedData, data) {
		panic(fmt.Errorf("supplied data did not match in call to Eval expected: %#v, actual: %#v", m.evalExpectedData, data))
	}

	m.evalCalled = true
	return m.evalResult, m.evalErr
}

func (m *MockNode) String() string {
	m.stringCalled = true
	return m.stringResult
}

func (m *MockNode) AssertEvalCalled(t *testing.T) {
	if !m.evalCalled {
		fmt.Println("expected call to Eval")
		t.Fail()
	}
}

func (m *MockNode) AssertStringCalled(t *testing.T) {
	if !m.stringCalled {
		fmt.Println("expected call to String")
		t.Fail()
	}
}

type MockNodeOp struct {
	stringCalled bool

	stringResult string
}

var _ NodeOp = &MockNodeOp{}

func NewMockNodeOp(stringResult string) *MockNodeOp {
	return &MockNodeOp{
		stringResult: stringResult,
	}
}

func (m *MockNodeOp) String() string {
	m.stringCalled = true
	return m.stringResult
}

func (m *MockNodeOp) AssertStringCalled(t *testing.T) {
	if !m.stringCalled {
		fmt.Println("expected call to String")
		t.Fail()
	}
}

type MockUnaryNodeOp struct {
	MockNodeOp

	calculateInput  any
	calculateResult any
	calculateErr    error

	calculateCalled bool
}

var _ UnaryNodeOp = &MockUnaryNodeOp{}

func NewMockUnaryNodeOp(
	calculateInput any,
	calculateResult any,
	calculateErr error,
	stringResult string,
) *MockUnaryNodeOp {
	return &MockUnaryNodeOp{
		MockNodeOp:      *NewMockNodeOp(stringResult),
		calculateInput:  calculateInput,
		calculateResult: calculateResult,
		calculateErr:    calculateErr,
	}
}

// Calculate implements UnaryNodeOp.
func (m *MockUnaryNodeOp) Calculate(right any) (any, error) {
	if !reflect.DeepEqual(m.calculateInput, right) {
		panic(fmt.Errorf("supplied data did not match in call to Calculate expected: %#v, actual: %#v", m.calculateInput, right))
	}

	m.calculateCalled = true
	return m.calculateResult, m.calculateErr
}

func (m *MockUnaryNodeOp) AssertCalculateCalled(t *testing.T) {
	if !m.stringCalled {
		fmt.Println("expected call to Calculate")
		t.Fail()
	}
}

type MockBinaryNodeOp struct {
	MockNodeOp

	calculateLeft   any
	calculateRight  any
	calculateResult any
	calculateErr    error

	calculateCalled bool
}

var _ BinaryNodeOp = &MockBinaryNodeOp{}

func NewMockBinaryNodeOp(
	calculateLeft any,
	calculateRight any,
	calculateResult any,
	calculateErr error,
	stringResult string,
) *MockBinaryNodeOp {
	return &MockBinaryNodeOp{
		MockNodeOp:      *NewMockNodeOp(stringResult),
		calculateLeft:   calculateLeft,
		calculateRight:  calculateRight,
		calculateResult: calculateResult,
		calculateErr:    calculateErr,
	}
}

// Calculate implements UnaryNodeOp.
func (m *MockBinaryNodeOp) Calculate(right, left any) (any, error) {
	if !reflect.DeepEqual(m.calculateLeft, left) {
		panic(fmt.Errorf("supplied data did not match in call to Calculate expected: %#v, actual: %#v", m.calculateLeft, right))
	}

	if !reflect.DeepEqual(m.calculateRight, right) {
		panic(fmt.Errorf("supplied data did not match in call to Calculate expected: %#v, actual: %#v", m.calculateRight, right))
	}

	m.calculateCalled = true
	return m.calculateResult, m.calculateErr
}

func (m *MockBinaryNodeOp) AssertCalculateCalled(t *testing.T) {
	if !m.stringCalled {
		fmt.Println("expected call to Calculate")
		t.Fail()
	}
}
