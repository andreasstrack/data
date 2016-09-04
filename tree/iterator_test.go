package tree

import (
	"fmt"
	"math"
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestBreadtFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	tree := NewTree(1)
	const lastValue int64 = 5
	for i := int64(2); i <= lastValue; i++ {
		tree.Add(NewTree(i))
	}
	c := tree.GetChildren()
	for i := range c {
		c[i].Add(NewTree(lastValue + int64(i) + 1))
	}

	ni := NewNodeIterator(tree, newIntChildIterator, BreadthFirst)

	for i := int64(0); i < lastValue*2-1; i++ {
		tt.Assert(ni.HasNext(), "HasNext (loop index %d)", i)
		tt.AssertEquals(ni.Next().(Node).GetValue().Int(), i+1, "Value (loop index %d)", i)
	}
	tt.Assert(!ni.HasNext(), "!HasNext (after loop): %s", ni.Next())
	tt.AssertEquals(ni.Next(), nil, "Value (after loop)")
}

func TestDepthFirstTraversal(t *testing.T) {
	// tt := T.NewT(t)
	// tree := NewTree(1)
	// const lastValue int64 = 20
	// bni := NewNodeIterator(tree, func(n Node) ChildIterator {
	// 	return newIntBuildingChildIterator(n, lastValue, 2)
	// }, BreadthFirst)
	// for bni.HasNext() {
	// 	next := bni.Next()
	// 	tt.Assert(next != nil, "next node: %s", next)
	// 	t.FailNow()
	// }
	// tt.AssertEquals(Size(tree), int(lastValue), "size of built tree")
	// fmt.Printf("%s\n", tree)
}

type intChildIterator struct {
	parent    Node
	nextIndex int
	next      Node
}

func (ici *intChildIterator) Init(n Node) {
	ici.parent = n
	ici.nextIndex = -1
	ici.getNext()
}

func (ici *intChildIterator) getNext() {
	c := ici.parent.GetChildren()
	l := len(c)
	ici.nextIndex = ici.nextIndex + 1
	if ici.nextIndex < l {
		ici.next = c[ici.nextIndex]
	} else {
		ici.next = nil
	}
}

func (ici *intChildIterator) HasNext() bool {
	return ici.next != nil
}

func (ici *intChildIterator) Next() interface{} {
	result := ici.next
	ici.getNext()
	return result
}

func newIntChildIterator(n Node) ChildIterator {
	ci := &intChildIterator{}
	ci.Init(n)
	return ci
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
	unclesNephewsAndSiblings := append(append(GetChildrenAndNephews(ibci.parent), GetSelfAndSiblings(ibci.parent)...), GetParentAndUncles(ibci.parent)...)
	fmt.Printf("Num uncles, nephews, and siblings: %d\n", len(unclesNephewsAndSiblings))
	if len(unclesNephewsAndSiblings) == 0 {
		return 0
	}
	for i := range unclesNephewsAndSiblings {
		maxValue = int64(math.Max(float64(maxValue), float64(unclesNephewsAndSiblings[i].GetValue().Int())))
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
