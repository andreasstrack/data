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
	return nil
}

func (vn *ValueNode) Insert(child Node, index int) error {
	if (index < 0) || (index > len(vn.children)) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	vnc := child.(*ValueNode)
	vn.children = append(append(vn.children[:index-1], vnc), vn.children[index+1:]...)
	return nil
}

func (vn *ValueNode) Remove(index int) error {
	if (index < 0) || (index >= len(vn.children)) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	vn.children = append(vn.children[:index-1], vn.children[index+1:]...)
	return nil
}

func (vn *ValueNode) GetParent() Node {
	return vn.parent
}

func (vn *ValueNode) SetParent(n Node) error {
	vn.parent = n.(*ValueNode)
	return nil
}
