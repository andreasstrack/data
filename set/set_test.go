package set

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestNewSetContentNotNil(t *testing.T) {
	tt := T.NewT(t)
	s, err := NewSet()
	tt.AssertNoError(err, "NewSet()")
	tt.Assert(s.content != nil, "Content of new set not nil.")
}

func TestInsert(t *testing.T) {
	tt := T.NewT(t)
	s, err := NewSet()
	tt.AssertNoError(err, "NewSet()")

	element1 := 3
	s.Insert(element1)
	tt.AssertEquals(1, s.Size(), "Size of set after first insert.")
	tt.Assert(s.Contains(element1), "Set contains first inserted element %#v: %#v", element1, s)

	element2 := "Hello"
	s.Insert(element2)
	tt.AssertEquals(2, s.Size(), "Size of set after second insert.")
	tt.Assert(s.Contains(element2), "Set contains second inserted element %#v: %#v", element2, s)

	s.Insert(element1)
	tt.AssertEquals(2, s.Size(), "Size of set after re-insert of first element.")
	tt.Assert(s.Contains(element1), "Set contains first inserted element %#v: %#v", element1, s)
}

func TestRemove(t *testing.T) {
	tt := T.NewT(t)
	s, err := NewSet()
	tt.AssertNoError(err, "NewSet()")

	element1 := 3
	s.Insert(element1)
	tt.Assert(s.Contains(element1), "Set contains first inserted element %#v: %#v", element1, s)

	s.Remove(element1)
	tt.Assert(s.IsEmpty(), "Set is empty after removing the only element.")
	tt.Assert(!s.Contains(element1), "Set does not contain first inserted element (%#v) anymore after removal: %#v", element1, s)
}

func TestSetFromSlice(t *testing.T) {
	tt := T.NewVerboseT(t)
	var slice []interface{}
	slice = append(slice, 1, "3", 5.3)
	s, err := NewSetFromSlice(&slice)
	tt.AssertNoError(err, "NewSetFromSlice()")
	tt.AssertEquals(len(slice), s.Size(), "Size of set is equal to the size of the slice the set was generated from (set: %#v).", s)
	for i := range slice {
		tt.Assert(s.Contains(slice[i]), "Set contains element %d: %#v", i, slice[i])
	}
}
