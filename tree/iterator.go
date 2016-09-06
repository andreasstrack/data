package tree

import (
	"github.com/andreasstrack/datastructures"
	"github.com/andreasstrack/util/patterns"
)

type ChildIterator interface {
	patterns.Iterator
	Init(n Node)
}

type ChildIteratorFactory func(n Node) ChildIterator

func newDefaultChildIterator(n Node) ChildIterator {
	ci := &defaultChildIterator{}
	ci.Init(n)
	return ci
}

type defaultChildIterator struct {
	parent    Node
	nextIndex int
	next      Node
}

func (ici *defaultChildIterator) Init(n Node) {
	ici.parent = n
	ici.nextIndex = -1
	ici.getNext()
}

func (ici *defaultChildIterator) getNext() {
	c := ici.parent.GetChildren()
	l := len(c)
	ici.nextIndex = ici.nextIndex + 1
	if ici.nextIndex < l {
		ici.next = c[ici.nextIndex]
	} else {
		ici.next = nil
	}
}

func (ici *defaultChildIterator) HasNext() bool {
	return ici.next != nil
}

func (ici *defaultChildIterator) Next() interface{} {
	result := ici.next
	ici.getNext()
	return result
}

type nodeIteratorStrategy interface {
	init(n Node, cbf ChildIteratorFactory)
	patterns.Iterator
}

type depthFirstStrategy struct {
	cbf  ChildIteratorFactory
	cbs  datastructures.Stack
	next Node
}

func (dfs *depthFirstStrategy) init(n Node, cbf ChildIteratorFactory) {
	dfs.cbf = cbf
	dfs.cbs.Clear()
	dfs.next = n
	dfs.newChildIterator()
}

func (dfs *depthFirstStrategy) newChildIterator() {
	dfs.cbs.Insert(dfs.cbf(dfs.next))
}

func (dfs *depthFirstStrategy) getNext() {
	if dfs.cbs.IsEmpty() {
		dfs.next = nil
		return
	}
	ci := dfs.cbs.Peek().(ChildIterator)
	if !ci.HasNext() {
		dfs.cbs.Pop()
		dfs.getNext()
		return
	}
	dfs.next = ci.Next().(Node)
	dfs.newChildIterator()
}

func (dfs *depthFirstStrategy) HasNext() bool {
	return dfs.next != nil
}

func (dfs *depthFirstStrategy) Next() interface{} {
	node := dfs.next
	dfs.getNext()
	return node
}

type breadthFirstStrategy struct {
	cif                ChildIteratorFactory
	ci                 ChildIterator
	currentParentNodes datastructures.Queue
	nextParentNodes    datastructures.Queue
	next               Node
}

func (bfs *breadthFirstStrategy) init(n Node, cif ChildIteratorFactory) {
	bfs.cif = cif
	bfs.next = n
	bfs.currentParentNodes = datastructures.NewFifoQueue()
	bfs.nextParentNodes = datastructures.NewFifoQueue()
	bfs.currentParentNodes.Insert(n)
	bfs.newChildIterator()
}

func (bfs *breadthFirstStrategy) newChildIterator() {
	if bfs.currentParentNodes.IsEmpty() {
		bfs.currentParentNodes = bfs.nextParentNodes
		bfs.nextParentNodes = datastructures.NewFifoQueue()
	}
	if bfs.currentParentNodes.IsEmpty() {
		bfs.ci = nil
		return
	}
	parentNode := bfs.currentParentNodes.Pop().(Node)
	bfs.ci = bfs.cif(parentNode)
}

func (bfs *breadthFirstStrategy) getNext() {
	if bfs.ci == nil {
		bfs.next = nil
		return
	}
	if !bfs.ci.HasNext() {
		bfs.newChildIterator()
		bfs.getNext()
		return
	}
	bfs.next = bfs.ci.Next().(Node)
	bfs.nextParentNodes.Insert(bfs.next)
}

func (bfs *breadthFirstStrategy) HasNext() bool {
	return bfs.next != nil
}

func (bfs *breadthFirstStrategy) Next() interface{} {
	node := bfs.next
	bfs.getNext()
	return node
}

type NodeIterator struct {
	nis  nodeIteratorStrategy
	next Node
}

func NewNodeIterator(start Node, cif ChildIteratorFactory, strategy TraversalStrategy) *NodeIterator {
	ni := &NodeIterator{}
	ni.init(start, cif, strategy)
	return ni
}

func (ni *NodeIterator) init(start Node, cif ChildIteratorFactory, strategy TraversalStrategy) {
	switch strategy {
	case DepthFirst:
		ni.nis = &depthFirstStrategy{}
	case BreadthFirst:
		ni.nis = &breadthFirstStrategy{}
	default:
		panic("invalid traversal strategy")
	}
	ni.nis.init(start, cif)
	if ni.nis.HasNext() {
		ni.next = ni.nis.Next().(Node)
	} else {
		ni.next = nil
	}
}

func (ni *NodeIterator) HasNext() bool {
	return ni.next != nil
}

func (ni *NodeIterator) Next() interface{} {
	result := ni.next
	if !ni.nis.HasNext() {
		ni.next = nil
	} else {
		ni.next = ni.nis.Next().(Node)
	}
	return result
}
