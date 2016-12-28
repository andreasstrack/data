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

func NewValueNode(i interface{}) *ValueNode {
	n := &ValueNode{}
	n.Value = reflect.ValueOf(i)
	n.parent = nil
	n.children = make([]Node, 0)
	return n
}

func (vn *ValueNode) String() string {
	switch vn.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%s", vn.Bool())
	case reflect.Int:
		return fmt.Sprintf("%d", vn.Int())
	case reflect.Uint:
		return fmt.Sprintf("%d", vn.Uint())
	case reflect.Float32:
		return fmt.Sprintf("%f", vn.Float())
	case reflect.Float64:
		return fmt.Sprintf("%f", vn.Float())
	case reflect.String:
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
	vnc := child.(*ValueNode)
	vn.children = append(vn.children, vnc)
	vnc.parent = vn
	return nil
}

func (vn *ValueNode) Insert(child Node, index int) error {
	if (index < 0) || (index > len(vn.children)) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	vnc := child.(*ValueNode)
	vn.children = append(append(vn.children[:index], vnc), vn.children[index:]...)
	vnc.parent = vn
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
	vn.parent = n.(*ValueNode)
	return nil
}
