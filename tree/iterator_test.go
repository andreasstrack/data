package tree

import (
	"fmt"
	"math"
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestBreadtFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()

	ni := NewNodeIterator(tree, newDefaultChildIterator, BreadthFirst)
	maxValue := int64(Size(tree))
	for i := int64(0); i < maxValue; i++ {
		tt.Assert(ni.HasNext(), "HasNext (loop index %d)", i)
		tt.AssertEquals(ni.Next().(Node).GetValue().Int(), i+1, "Value (loop index %d)", i)
	}
	tt.Assert(!ni.HasNext(), "!HasNext (after loop): %s", ni.Next())
	tt.AssertEquals(ni.Next(), nil, "Value (after loop)")
}

func buildIntTree(lastValue int64) *Tree {
	tree := NewTree(1)
	bni := NewNodeIterator(tree, func(n Node) ChildIterator {
		return newIntBuildingChildIterator(n, lastValue, 2)
	}, BreadthFirst)
	for bni.HasNext() {
		bni.Next()
	}
	return tree
}

func TestIntBuildingChildIterator(t *testing.T) {
	tt := T.NewT(t)
	const lastValue int64 = 7
	tree := buildIntTree(lastValue)
	tt.Assert(tree != nil, "resulting tree not nil: %s", tree)
	tt.AssertEquals(Size(tree), int(lastValue), "size of built tree")
	tt.AssertEquals("1 -> [2 -> [4 5] 3 -> [6 7]]", tree.String(), "string representation of tree")
}

func TestDepthFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	const lastValue int64 = 7
	tree := buildIntTree(lastValue)
	ni := NewNodeIterator(tree, newDefaultChildIterator, DepthFirst)
	var expectedNodeValues = [...]int64{1, 2, 4, 5, 3, 6, 7}
	index := 0
	for ni.HasNext() {
		n := ni.Next().(Node)
		tt.AssertEquals(expectedNodeValues[index], n.GetValue().Int(), "value of index %d in %s", index, tree)
		index++
	}
	tt.AssertEquals(index, Size(tree), "number of iterated nodes in %s", tree)
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
	fmt.Printf("Num uncles, nephews, and siblings: %d\n", len(allNodes))
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
	ibci.next = NewTree(ibci.nextValue)
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
