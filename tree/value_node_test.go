package tree

import (
	"testing"

	"github.com/andreasstrack/datastructures"
	T "github.com/andreasstrack/util/testing"
)

func TestValueNodeAsValue(t *testing.T) {
	tt := T.NewT(t)
	vn := NewValueNodeFromInterface(1)
	var v datastructures.Value
	v = vn
	tt.Assert(v.IsInt(), "value type (%s) is Int\n", v)
}
