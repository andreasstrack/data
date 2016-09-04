package datastructures

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestArrayList(t *testing.T) {
	tt := T.NewT(t)
	l := NewArrayList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	tt.AssertEquals(l.Size(), 3, "size of %s", l)
	tt.AssertEquals(l.Front(), 1, "front of %s", l)
	tt.AssertEquals(l.Back(), 3, "back of %s", l)
	tt.AssertEquals(l.At(1), 2, "at(1) of %s", l)
	l.Remove(1)
	tt.AssertEquals(l.Size(), 2, "size of %s", l)
	tt.AssertEquals(l.At(1), 3, "at(1) of %s", l)
	tt.Assert(!l.IsEmpty(), "IsEmpty (%s)?", l)
	l.Clear()
	tt.Assert(l.IsEmpty(), "IsEmpty (%s)?", l)
}
