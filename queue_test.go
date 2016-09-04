package datastructures

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestFifoQueue(t *testing.T) {
	s := NewFifoQueue()
	tt := T.NewT(t)

	tt.AssertEquals(s.Size(), 0, "Size of %s", s)
	tt.Assert(s.IsEmpty(), "IsEmpty? %s")

	s.Insert(4)
	tt.AssertEquals(s.Size(), 1, "Size of %s", s)
	tt.AssertEquals(s.Peek().(int), 4, "Peek of %s", s)
	s.Insert(5)
	s.Insert(10)
	tt.AssertEquals(s.Size(), 3, "Size of %s", s)
	tt.AssertEquals(s.Pop(), 4, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), 5, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), 10, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), nil, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), nil, "Popped top of %s", s)
	s.Insert(11)
	tt.AssertEquals(s.Size(), 1, "Size of %s", s)
	s.Clear()
	tt.AssertEquals(s.Size(), 0, "Size of %s", s)
	tt.Assert(s.IsEmpty(), "IsEmpty? %s")
}

func TestStack(t *testing.T) {
	s := NewStack()
	tt := T.NewT(t)

	tt.AssertEquals(s.Size(), 0, "Size of %s", s)
	tt.Assert(s.IsEmpty(), "IsEmpty? %s")

	s.Insert(4)
	tt.AssertEquals(s.Size(), 1, "Size of %s", s)
	tt.AssertEquals(s.Peek().(int), 4, "Peek of %s", s)
	s.Insert(5)
	s.Insert(10)
	tt.AssertEquals(s.Size(), 3, "Size of %s", s)
	tt.AssertEquals(s.Pop(), 10, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), 5, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), 4, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), nil, "Popped top of %s", s)
	tt.AssertEquals(s.Pop(), nil, "Popped top of %s", s)
	s.Insert(11)
	tt.AssertEquals(s.Size(), 1, "Size of %s", s)
	s.Clear()
	tt.AssertEquals(s.Size(), 0, "Size of %s", s)
	tt.Assert(s.IsEmpty(), "IsEmpty? %s")
}
