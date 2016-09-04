package tree

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestString(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	const referenceString = "1 -> [2 -> [4] 3 -> [5 6]]"
	tt.AssertEquals(referenceString, tree.String(), "string representation of tree")
}

func TestSize(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	tt.AssertEquals(6, Size(tree), "size of tree '%s'", tree)
}

func TestDepth(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	tt.AssertEquals(3, Depth(tree), "depth of tree '%s'", tree)
}

func TestBranchingFactor(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	tt.AssertEquals(2, BranchingFactor(tree), "branching factor of tree '%s'", tree)
}

func TestGetParentAndUncles(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	uncles := GetParentAndUncles(tree.GetChildren()[0].GetChildren()[0])
	tt.AssertEquals(2, len(uncles), "amount of parent and uncles of '4' in '%s'", tree)
	tt.AssertEquals(int64(2), uncles[0].GetValue().Int(), "first uncle")
	tt.AssertEquals(int64(3), uncles[1].GetValue().Int(), "second uncle")
}

func TestGetSelfAndSiblings(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	siblings := GetSelfAndSiblings(tree.GetChildren()[0])
	tt.AssertEquals(2, len(siblings), "amount of self and siblings of '2' in '%s'", tree)
	tt.AssertEquals(int64(2), siblings[0].GetValue().Int(), "first sibling")
	tt.AssertEquals(int64(3), siblings[1].GetValue().Int(), "second sibling")
}

func TestGetSelfAndCousins(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	cousins := GetSelfSiblingsAndCousins(tree.GetChildren()[0].GetChildren()[0])
	tt.AssertEquals(3, len(cousins), "amount of self, siblings, and cousins of '4' in '%s'", tree)
	tt.AssertEquals(int64(4), cousins[0].GetValue().Int(), "first sibling")
	tt.AssertEquals(int64(5), cousins[1].GetValue().Int(), "second sibling")
	tt.AssertEquals(int64(6), cousins[2].GetValue().Int(), "third sibling")
}

func TestGetChildrenAndNephews(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	nephews := GetChildrenAndNephews(tree.GetChildren()[0])
	tt.AssertEquals(3, len(nephews), "amount of children and nephews of '2' in '%s'", tree)
	tt.AssertEquals(int64(4), nephews[0].GetValue().Int(), "first child/nephew")
	tt.AssertEquals(int64(5), nephews[1].GetValue().Int(), "second child/nephew")
	tt.AssertEquals(int64(6), nephews[2].GetValue().Int(), "third child/nephew")
}

func buildTestTree() *Tree {
	tree := NewTree(1)
	tree.Add(NewTree(2))
	tree.Add(NewTree(3))
	tree.GetChildren()[0].Add(NewTree(4))
	tree.GetChildren()[1].Add(NewTree(5))
	tree.GetChildren()[1].Add(NewTree(6))
	return tree
}
