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

func TestGetParentAndUncles(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	{
		uncles := GetParentAndUncles(tree)
		tt.AssertEquals(0, len(uncles), "amount of parent and uncles of '1' in '%s'", String(tree))
	}
	{
		uncles := GetParentAndUncles(tree.GetChildren()[0])
		tt.AssertEquals(1, len(uncles), "amount of parent and uncles of '2' in '%s'", String(tree))
		tt.AssertEquals(int64(1), uncles[0].GetValue().Int(), "first uncle")
	}
	{
		uncles := GetParentAndUncles(tree.GetChildren()[0].GetChildren()[0])
		tt.AssertEquals(2, len(uncles), "amount of parent and uncles of '4' in '%s'", String(tree))
		tt.AssertEquals(int64(2), uncles[0].GetValue().Int(), "first uncle")
		tt.AssertEquals(int64(3), uncles[1].GetValue().Int(), "second uncle")
	}
}

func TestGetSelfAndSiblings(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	{
		siblings := GetSelfAndSiblings(tree)
		tt.AssertEquals(1, len(siblings), "amount of self and siblings of '1' in '%s'", String(tree))
		tt.AssertEquals(int64(1), siblings[0].GetValue().Int(), "first sibling")
	}
	{
		siblings := GetSelfAndSiblings(tree.GetChildren()[0])
		tt.AssertEquals(2, len(siblings), "amount of self and siblings of '2' in '%s'", String(tree))
		tt.AssertEquals(int64(2), siblings[0].GetValue().Int(), "first sibling")
		tt.AssertEquals(int64(3), siblings[1].GetValue().Int(), "second sibling")
	}
}

func TestGetSelfSiblingsAndCousins(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	{
		cousins := GetSelfSiblingsAndCousins(tree)
		tt.AssertEquals(1, len(cousins), "amount of self, siblings, and cousins of '1' in '%s'", String(tree))
		tt.AssertEquals(int64(1), cousins[0].GetValue().Int(), "first cousin")
	}
	{
		cousins := GetSelfSiblingsAndCousins(tree.GetChildren()[0])
		tt.AssertEquals(2, len(cousins), "amount of self, siblings, and cousins of '2' in '%s'", String(tree))
		tt.AssertEquals(int64(2), cousins[0].GetValue().Int(), "first cousin")
		tt.AssertEquals(int64(3), cousins[1].GetValue().Int(), "second cousin")
	}
	{
		cousins := GetSelfSiblingsAndCousins(tree.GetChildren()[0].GetChildren()[0])
		tt.AssertEquals(3, len(cousins), "amount of self, siblings, and cousins of '4' in '%s'", String(tree))
		tt.AssertEquals(int64(4), cousins[0].GetValue().Int(), "first cousin")
		tt.AssertEquals(int64(5), cousins[1].GetValue().Int(), "second cousin")
		tt.AssertEquals(int64(6), cousins[2].GetValue().Int(), "third cousin")
	}
}

func TestGetChildrenAndNephews(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	{
		nephews := GetChildrenAndNephews(tree)
		tt.AssertEquals(2, len(nephews), "amount of children and nephews of '1' in '%s'", String(tree))
		tt.AssertEquals(int64(2), nephews[0].GetValue().Int(), "first child/nephew")
		tt.AssertEquals(int64(3), nephews[1].GetValue().Int(), "second child/nephew")
	}
	{
		nephews := GetChildrenAndNephews(tree.GetChildren()[0])
		tt.AssertEquals(3, len(nephews), "amount of children and nephews of '2' in '%s'", String(tree))
		tt.AssertEquals(int64(4), nephews[0].GetValue().Int(), "first child/nephew")
		tt.AssertEquals(int64(5), nephews[1].GetValue().Int(), "second child/nephew")
		tt.AssertEquals(int64(6), nephews[2].GetValue().Int(), "third child/nephew")
	}
	nephews := GetChildrenAndNephews(tree.GetChildren()[0])
	tt.AssertEquals(3, len(nephews), "amount of children and nephews of '2' in '%s'", String(tree))
	tt.AssertEquals(int64(4), nephews[0].GetValue().Int(), "first child/nephew")
	tt.AssertEquals(int64(5), nephews[1].GetValue().Int(), "second child/nephew")
	tt.AssertEquals(int64(6), nephews[2].GetValue().Int(), "third child/nephew")
}

func TestGetRoot(t *testing.T) {
	tt := T.NewT(t)
	tree := buildTestTree()
	{
		root := GetRoot(tree)
		tt.AssertEquals(tree, root, "root of %s", tree)
	}
	{
		n := tree.GetChildren()[1].GetChildren()[0]
		root := GetRoot(n)
		tt.AssertEquals(tree, root, "root of %s in %s", n, tree)
	}
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
