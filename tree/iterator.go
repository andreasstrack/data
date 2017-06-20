package tree

import (
	"github.com/andreasstrack/data"
	"github.com/andreasstrack/util/patterns"
)

type NodeValidator interface {
	IsValid(n Node) bool
}

type defaultNodeValidator struct {
}

func (dnv defaultNodeValidator) IsValid(n Node) bool {
	return true
}

type NodeIterator struct {
	nis  nodeIteratorStrategy
	nv   NodeValidator
	next Node
}

func NewValidatedNodeIterator(start Node, cif ChildIteratorFactory, strategy TraversalStrategy, nv NodeValidator) *NodeIterator {
	ni := &NodeIterator{}
	ni.init(start, cif, strategy, nv)
	return ni
}

func NewNodeIterator(start Node, cif ChildIteratorFactory, strategy TraversalStrategy) *NodeIterator {
	ni := &NodeIterator{}
	ni.init(start, cif, strategy, defaultNodeValidator{})
	return ni
}

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
	cbs  data.Stack
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
	cif         ChildIteratorFactory
	ci          ChildIterator
	nextParents data.Queue
	next        Node
}

func (bfs *breadthFirstStrategy) init(n Node, cif ChildIteratorFactory) {
	bfs.cif = cif
	bfs.next = n
	bfs.nextParents = data.NewFifoQueue()
	bfs.nextParents.Insert(n)
	bfs.newChildIterator()
}

func (bfs *breadthFirstStrategy) newChildIterator() {
	if bfs.nextParents.IsEmpty() {
		bfs.ci = nil
		return
	}
	parentNode := bfs.nextParents.Pop().(Node)
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
	bfs.nextParents.Insert(bfs.next)
}

func (bfs *breadthFirstStrategy) HasNext() bool {
	return bfs.next != nil
}

func (bfs *breadthFirstStrategy) Next() interface{} {
	node := bfs.next
	bfs.getNext()
	return node
}

func (ni *NodeIterator) init(start Node, cif ChildIteratorFactory, strategy TraversalStrategy, nv NodeValidator) {
	ni.nv = nv
	switch strategy {
	case DepthFirst:
		ni.nis = &depthFirstStrategy{}
	case BreadthFirst:
		ni.nis = &breadthFirstStrategy{}
	default:
		panic("invalid traversal strategy")
	}
	ni.nis.init(start, cif)
	ni.next = nil
	_ = ni.Next()
}

func (ni *NodeIterator) HasNext() bool {
	return ni.next != nil
}

func (ni *NodeIterator) Next() interface{} {
	var result Node
	resultOk := false
	nextOk := false
	for !resultOk || !nextOk {
		if !resultOk {
			result = ni.next
		}
		if !ni.nis.HasNext() {
			ni.next = nil
		} else {
			ni.next = ni.nis.Next().(Node)
		}
		resultOk = result == nil || ni.nv.IsValid(result)
		nextOk = ni.next == nil || ni.nv.IsValid(ni.next)
	}
	return result
}
