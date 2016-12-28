package tree

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestValueNode(t *testing.T) {
	tt := T.NewT(t)
	root := NewValueNode(1)
	for i := 2; i < 5; i++ {
		tt.AssertNoError(root.Add(NewValueNode(i)), "add node %d to roor", i)
	}

}
