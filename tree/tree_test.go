package tree

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestGetSetParent(t *testing.T) {
	tt := T.NewT(t)
	tree := NewValueNodeFromInterface(5)

	parentTree := NewValueNodeFromInterface(10)

	tree.SetParent(parentTree)
	tt.AssertEquals(tree.GetParent(), parentTree, "parent of %s", tree)
}

func TestAddChildren(t *testing.T) {
	tt := T.NewT(t)
	tree := NewValueNodeFromInterface(1)
	tree.Add(NewValueNodeFromInterface(2))
	tree.Add(NewValueNodeFromInterface(3))
	tree.Add(NewValueNodeFromInterface(4))
	tt.AssertEquals(tree.GetChildren()[0].GetValue().Int(), int64(2), "value of child 0", tree)
	tt.AssertEquals(tree.GetChildren()[1].GetValue().Int(), int64(3), "value of child 1", tree)
	tt.AssertEquals(tree.GetChildren()[2].GetValue().Int(), int64(4), "value of child 2", tree)
}

func TestInsertRemoveChildren(t *testing.T) {
	tt := T.NewT(t)
	tree := NewValueNodeFromInterface(1)
	tt.AssertError(tree.Insert(NewValueNodeFromInterface(2), 1), "insert at invalid index")
	tt.AssertError(tree.Insert(NewValueNodeFromInterface(2), -1), "insert at invalid index")
	tt.AssertNoError(tree.Insert(NewValueNodeFromInterface(2), 0), "insert at valid index")
	tt.AssertNoError(tree.Insert(NewValueNodeFromInterface(3), 1), "insert at valid index")
	tt.AssertNoError(tree.Insert(NewValueNodeFromInterface(4), 1), "insert at valid index")
	tt.AssertNoError(tree.Remove(1), "Remove at valid index")
	tt.AssertError(tree.Remove(2), "remove at invalid index")
	tt.AssertEquals(tree.GetChildren()[0].GetValue().Int(), int64(2), "value of child 0")
	tt.AssertEquals(tree.GetChildren()[1].GetValue().Int(), int64(4), "value of child 1")
	tt.AssertNoError(tree.Remove(1), "Remove at valid index")
	tt.AssertNoError(tree.Remove(0), "Remove at valid index")
	tt.AssertEquals(len(tree.GetChildren()), 0, "number of children of %s", tree)
}
