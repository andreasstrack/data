package datastructures

import (
	"fmt"
	"math/rand"
)

// The Comparable interface declares a couple of
// comparison functions useful for, for example, sorting.
//
// NOTE: The use of the interface as parameters to the
// functions does not imply that the function parameters
// are of the same type as the callee. However, I have not
// found another way to model this in Go (until now), and at
// least it implies that types implementing this interface
// are comparable to themselves, which is probably the most
// widespread use case anyway.
type Comparable interface {
	// LessThan performs the 'less' comparison
	// to another Comparable.
	LessThan(c Comparable) bool

	// EqualTo performs the 'equal' comparison
	// to another Comparable.
	EqualTo(c Comparable) bool
}

// A ComparableList is a slice of Comparable objects.
type ComparableList []Comparable

// IsSorted returns whether a ComparableList is sorted.
func (cl ComparableList) IsSorted() bool {
	l := len(cl)
	for i := 0; i < l-1; i++ {
		if cl[i+1].LessThan(cl[i]) {
			return false
		}
	}
	return true
}

// Swap swaps the elements with i and j.
func (cl *ComparableList) Swap(i, j int) {
	jOld := (*cl)[j]
	(*cl)[j] = (*cl)[i]
	(*cl)[i] = jOld
}

// Split returns two ComparableLists resulting from a split of cl.
// The index i indicates the start element of the second resulting list.
// TODO: Use uint as parameter? i < 0 will now cause a panic.
func (cl *ComparableList) Split(i int) (*ComparableList, *ComparableList) {
	cll := (*cl)[0:i]
	clr := (*cl)[i:]
	return &cll, &clr
}

// SplitToHalves splits cl in two halves.
func (cl *ComparableList) SplitToHalves() (*ComparableList, *ComparableList) {
	if cl == nil {
		return nil, nil
	}

	return cl.Split(len(*cl) / 2)
}

// ComparableInt implements the Comparable interface
// for int data.
type ComparableInt int

func (ci ComparableInt) String() string {
	return fmt.Sprintf("%d", int(ci))
}

// LessThan implements the LessThan function for ComparableInt.
func (ci ComparableInt) LessThan(oci Comparable) bool {
	return ci < oci.(ComparableInt)
}

// EqualTo implements the EqualTo function for ComparableInt.
func (ci ComparableInt) EqualTo(oci Comparable) bool {
	return ci == oci.(ComparableInt)
}

// ComparableIntList is a slice of ComparableInt values.
type ComparableIntList []ComparableInt

// SortedIntList returns a sorted list of ComparableInt
// of length n.
func SortedIntList(n int) *ComparableList {
	var result ComparableList
	for i := 1; i <= n; i++ {
		result = append(result, ComparableInt(i))
	}
	return &result
}

// RandomIntList returns a list of random ComparableInt in
// the range of 0..upperLimit of length n.
func RandomIntList(n, upperLimit int) *ComparableList {
	var result ComparableList
	rand.Seed(1234)
	for i := 0; i < n; i++ {
		result = append(result, ComparableInt(rand.Intn(upperLimit)))
	}
	return &result
}

// ReverseSortedIntList returns a reverse sorted list of ComparableInt
// of length n.
func ReverseSortedIntList(n int) *ComparableList {
	var result ComparableList
	for i := n; i > 0; i-- {
		result = append(result, ComparableInt(i))
	}
	return &result
}
