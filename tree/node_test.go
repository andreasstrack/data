package tree

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestString(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	const referenceString = "1 -> [2 -> [4 -> [7 8]] 3 -> [5 -> [9 10] 6]]"
	tt.AssertEquals(referenceString, String(tree), "string representation of tree")
}

func TestSize(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	tt.AssertEquals(10, Size(tree), "size of tree '%s'", tree)
}

func TestDepth(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	tt.AssertEquals(4, Depth(tree), "depth of tree '%s'", tree)
}

func TestBranchingFactor(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	tt.AssertEquals(2, BranchingFactor(tree), "branching factor of tree '%s'", tree)
}

func buildTestTree() Node {
	tree := NewValueNodeFromInterface(1)
	tree.Add(NewValueNodeFromInterface(2))
	tree.Add(NewValueNodeFromInterface(3))
	tree.GetChildren()[0].Add(NewValueNodeFromInterface(4))
	tree.GetChildren()[1].Add(NewValueNodeFromInterface(5))
	tree.GetChildren()[1].Add(NewValueNodeFromInterface(6))
	tree.GetChildren()[0].GetChildren()[0].Add(NewValueNodeFromInterface(7))
	tree.GetChildren()[0].GetChildren()[0].Add(NewValueNodeFromInterface(8))
	tree.GetChildren()[1].GetChildren()[0].Add(NewValueNodeFromInterface(9))
	tree.GetChildren()[1].GetChildren()[0].Add(NewValueNodeFromInterface(10))
	return tree
}
