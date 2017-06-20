package tree

import (
	"testing"

	"github.com/andreasstrack/data"
	T "github.com/andreasstrack/util/testing"
)

func TestValueNodeAsValue(t *testing.T) {
	tt := T.NewT(t)
	vn := NewValueNodeFromInterface(1)
	var v data.Value
	v = vn
	tt.Assert(v.IsInt(), "value type (%s) is Int\n", v)
}
