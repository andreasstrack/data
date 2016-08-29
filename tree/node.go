package tree

import "github.com/andreasstrack/datastructures"

// A Node is a parent and a child,
// and it has a value.
type Node interface {
	datastructures.Child
	datastructures.Parent
	datastructures.Value
}
