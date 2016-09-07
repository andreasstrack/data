package tree

import (
	"fmt"

	"github.com/andreasstrack/datastructures"
)

type TraversalStrategy int

const (
	DepthFirst TraversalStrategy = iota
	BreadthFirst
)

type variantNode struct {
	val      datastructures.Variant
	parent   Node
	children []Node
}

func NewNode(value interface{}) Node {
	return &variantNode{val: *datastructures.NewVariant(value), parent: nil, children: make([]Node, 0)}
}

func (t *variantNode) GetValue() datastructures.Value {
	return &t.val
}

func (t *variantNode) GetChildren() (children []Node) {
	return t.children
}

func (t *variantNode) Add(child Node) {
	t.children = append(t.children, child)
	child.SetParent(t)
}

func (t *variantNode) Insert(child Node, index int) error {
	if index < 0 || index > len(t.children) {
		return fmt.Errorf("index %d out of range (0-%d)", index, len(t.children))
	}
	t.children = append(append(t.children[:index], child), t.children[index:]...)
	child.SetParent(t)
	return nil
}

func (t *variantNode) Remove(index int) error {
	if index < 0 || index >= len(t.children) {
		return fmt.Errorf("index %d out of range (0-%d)", index, len(t.children)-1)
	}
	t.children = append(t.children[0:index], t.children[index+1:]...)
	return nil
}

func (t *variantNode) GetParent() Node {
	return t.parent
}

func (t *variantNode) SetParent(parent Node) {
	t.parent = parent
}

func (t *variantNode) String() string {
	return String(t)
}
