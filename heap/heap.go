package heap

import (
	"github.com/andreasstrack/datastructures"
	"github.com/andreasstrack/datastructures/tree"
)

// A Heap is a tree, where all parents compare in the
// same way to their children, i.e. their values are either less
// (Min-Heap) or greater (Max-Heap) than the values of
// their children.
type Heap interface {
	tree.Node
	datastructures.PriorityQueue
}

// A BinaryHeap is a Heap, where each node has at most
// two children.
type BinaryHeap interface {
	tree.BinaryNode
	datastructures.PriorityQueue
}
