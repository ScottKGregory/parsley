// Package nodes provides implementation and interfaces for the various supported node types
package nodes

// NodeOp defines the basic capabilities of all operations
type NodeOp interface {
	String() string
}

// Node defines the basic capabilities of all nodes
type Node interface {
	Eval(data map[string]any) (any, error)
	String() string
}
