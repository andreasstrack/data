package datastructures

// A Queue is a collection to which
// elements can be added and from which
// elements can be retrieved. Upon retrieval,
// elements will be removed from the queue.
// In a queue the next retrievable element is
// identified according to the type / implementation
// of the queue, and this element is called the "head"
// of the queue. Upon retrieving an element from a
// queue, it will be removed from the queue.
type Queue interface {
	Add(e interface{})
	Remove()
	Element() interface{}
}

// A PriorityQueue is a queue where at any
// time the head is the element in the queue
// with the highest priority. The relative
// priority of elements is determined by comparing
// them.
type PriorityQueue interface {
	Add(c Comparable)
	Remove()
	Element() Comparable
}
