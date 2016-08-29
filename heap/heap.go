package heap

import "github.com/andreasstrack/datastructures"

// A Heap is a tree in which all parents compare in the
// same way to their children, i.e. their values are either less
// (Min-Heap) or greater (Max-Heap) than the values of
// their children.
type Heap interface {
	datastructures.PriorityQueue
}
