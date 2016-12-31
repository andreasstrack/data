package tree

import (
	"math"
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestDefaultChildIterator(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	ci := newDefaultChildIterator(tree)
	maxValue := int64(len(tree.GetChildren()))
	for i := int64(0); i < maxValue; i++ {
		tt.Assert(ci.HasNext(), "HasNext (loop index %d)", i)
		tt.AssertEquals(i+2, ci.Next().(Node).GetValue().Int(), "Value (loop index %d)", i)
	}
	tt.Assert(!ci.HasNext(), "!HasNext (after loop): %s", ci.Next())
	tt.AssertEquals(nil, ci.Next(), "Value (after loop)")
}

func TestBreadthFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()

	ni := NewNodeIterator(tree, newDefaultChildIterator, BreadthFirst)
	maxValue := int64(Size(tree))
	for i := int64(0); i < maxValue; i++ {
		tt.Assert(ni.HasNext(), "HasNext (loop index %d)", i)
		tt.AssertEquals(i+1, ni.Next().(Node).GetValue().Int(), "Value (loop index %d)", i)
	}
	tt.Assert(!ni.HasNext(), "!HasNext (after loop): %s", ni.Next())
	tt.AssertEquals(nil, ni.Next(), "Value (after loop)")
}

func TestDepthFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	const lastValue int64 = 7
	tree := buildIntTree(lastValue)
	ni := NewNodeIterator(tree, newDefaultChildIterator, DepthFirst)
	var expectedNodeValues = [...]int64{1, 2, 4, 5, 3, 6, 7}
	var index int
	for index = 0; ni.HasNext(); index++ {
		n := ni.Next().(Node)
		tt.AssertEquals(expectedNodeValues[index], n.GetValue().Int(), "value of index %d in %s", index) //, String(tree))
	}
	tt.AssertEquals(index, Size(tree), "number of iterated nodes in %s", tree)
}

func TestIntBuildingChildIterator(t *testing.T) {
	tt := T.NewT(t)
	const lastValue int64 = 7
	tree := buildIntTree(lastValue)
	tt.Assert(tree != nil, "resulting tree not nil: %s", String(tree))
	tt.AssertEquals(Size(tree), int(lastValue), "size of built tree")
	tt.AssertEquals("1 -> [2 -> [4 5] 3 -> [6 7]]", String(tree), "string representation of tree")
}

func buildIntTree(lastValue int64) Node {
	tree := NewValueNodeFromInterface(1)
	bni := NewNodeIterator(tree, func(n Node) ChildIterator {
		return newIntBuildingChildIterator(n, lastValue, 2)
	}, BreadthFirst)
	for bni.HasNext() {
		bni.Next()
	}
	return tree
}

type intBuildingChildIterator struct {
	parent          Node
	count           uint64
	nextValue       int64
	maxValue        int64
	branchingFactor uint64
	next            Node
}

func newIntBuildingChildIterator(n Node, maxValue int64, branchingFactor uint64) *intBuildingChildIterator {
	ibci := &intBuildingChildIterator{maxValue: maxValue, branchingFactor: branchingFactor}
	ibci.Init(n)
	return ibci
}

func (ibci *intBuildingChildIterator) Init(n Node) {
	ibci.count = 0
	ibci.parent = n
	ibci.nextValue = ibci.findNextValue()
	ibci.getNext()
}

func (ibci *intBuildingChildIterator) findNextValue() int64 {
	maxValue := int64(math.MinInt64)
	allNodes := GetAllNodesOfTree(ibci.parent)
	if len(allNodes) == 0 {
		return 0
	}
	for i := range allNodes {
		maxValue = int64(math.Max(float64(maxValue), float64(allNodes[i].GetValue().Int())))
	}
	return maxValue + 1
}

func (ibci *intBuildingChildIterator) getNext() {
	if ibci.count >= ibci.branchingFactor || ibci.nextValue > ibci.maxValue {
		ibci.next = nil
		return
	}
	ibci.next = NewValueNodeFromInterface(ibci.nextValue)
	ibci.nextValue++
	ibci.count++
}

func (ibci *intBuildingChildIterator) HasNext() bool {
	return ibci.next != nil
}

func (ibci *intBuildingChildIterator) Next() interface{} {
	if ibci.next == nil {
		return nil
	}
	ibci.parent.Add(ibci.next)
	result := ibci.next
	ibci.getNext()
	return result
}
