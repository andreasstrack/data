package datastructures

import (
	"fmt"
	"testing"

	T "github.com/andreasstrack/util/testing"
)

const numElements int = 02

func TestSplit(t *testing.T) {
	cl := RandomIntList(numElements, 1000)

	for i := 0; i < numElements; i++ {
		cll, clr := cl.Split(i)
		T.Assert(len(*cll) == i, "Expected left length: %d   Actual left length: %d\n", t, i, len(*cll))
		T.Assert(len(*clr) == numElements-i, "Expected right length: %d   Actual right length: %d\n", t, numElements-i, len(*clr))
	}
}

func TestSplitToHalves(t *testing.T) {
	cl := RandomIntList(numElements, 1000)

	cll, clr := cl.SplitToHalves()

	fmt.Printf("TestSplit:\n%s\n->\n%s\n,\n%s\n", *cl, *cll, *clr)

	splitIndex := numElements / 2
	for i := 0; i < splitIndex; i++ {
		if !(*cl)[i].EqualTo((*cll)[i]) {
			t.Errorf("Difference in element %d.", i)
		}
	}
	for i := splitIndex; i < numElements; i++ {
		if !(*cl)[i].EqualTo((*clr)[i-splitIndex]) {
			t.Errorf("Difference in element %d.", i)
		}
	}
}
