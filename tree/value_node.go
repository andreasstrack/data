package tree

import (
	"fmt"
	"reflect"

	"github.com/andreasstrack/datastructures"
)

type ValueNode struct {
	reflect.Value
	parent   Node
	children []Node
}

func NewValueNodeFromInterface(i interface{}) *ValueNode {
	return NewValueNode(reflect.ValueOf(i))
}

func NewValueNode(v reflect.Value) *ValueNode {
	n := &ValueNode{}
	n.Value = v
	n.parent = nil
	n.children = make([]Node, 0)
	return n
}

func (vn *ValueNode) IsBool() bool {
	return vn.Kind() == reflect.Bool
}

func (vn *ValueNode) IsInt() bool {
	switch vn.Kind() {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	default:
		return false
	}
}

func (vn *ValueNode) IsUint() bool {
	switch vn.Kind() {
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	default:
		return false
	}
}

func (vn *ValueNode) IsFloat() bool {
	switch vn.Kind() {
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	default:
		return false
	}
}

func (vn *ValueNode) IsString() bool {
	return vn.Kind() == reflect.String
}

func (vn *ValueNode) String() string {
	switch {
	case vn.IsBool():
		return fmt.Sprintf("%s", vn.Bool())
	case vn.IsInt():
		return fmt.Sprintf("%d", vn.Int())
	case vn.IsUint():
		return fmt.Sprintf("%d", vn.Uint())
	case vn.IsFloat():
		return fmt.Sprintf("%f", vn.Float())
	case vn.IsString():
		return vn.Value.String()
	default:
		return vn.Value.String()
	}
}

func (vn *ValueNode) GetValue() datastructures.Value {
	return vn
}

func (vn *ValueNode) ReflectValue() *reflect.Value {
	return &vn.Value
}

func (vn *ValueNode) GetChildren() []Node {
	return vn.children
}

func (vn *ValueNode) Add(child Node) error {
	vn.children = append(vn.children, child)
	child.SetParent(vn)
	return nil
}

func (vn *ValueNode) Insert(child Node, index int) error {
	if (index < 0) || (index > len(vn.children)) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	vn.children = append(append(vn.children[:index], child), vn.children[index:]...)
	child.SetParent(vn)
	return nil
}

func (vn *ValueNode) Remove(index int) error {
	if (index < 0) || (index >= len(vn.children)) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	vn.children = append(vn.children[:index], vn.children[index+1:]...)
	return nil
}

func (vn *ValueNode) GetParent() Node {
	return vn.parent
}

func (vn *ValueNode) SetParent(n Node) error {
	vn.parent = n
	return nil
}
