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

type Tree struct {
	val datastructures.Variant

	parent   Node
	children []Node
}

func NewTree(rootValue interface{}) *Tree {
	return &Tree{val: *datastructures.NewVariant(rootValue), parent: nil, children: make([]Node, 0)}
}

func (t *Tree) GetValue() datastructures.Value {
	return &t.val
}

func (t *Tree) GetChildren() (children []Node) {
	return t.children
}

func (t *Tree) Add(child Node) {
	t.children = append(t.children, child)
	child.SetParent(t)
}

func (t *Tree) Insert(child Node, index int) error {
	if index < 0 || index > len(t.children) {
		return fmt.Errorf("index %d out of range (0-%d)", index, len(t.children))
	}
	t.children = append(append(t.children[:index], child), t.children[index:]...)
	child.SetParent(t)
	return nil
}

func (t *Tree) Remove(index int) error {
	if index < 0 || index >= len(t.children) {
		return fmt.Errorf("index %d out of range (0-%d)", index, len(t.children)-1)
	}
	t.children = append(t.children[0:index], t.children[index+1:]...)
	return nil
}

func (t *Tree) GetParent() Node {
	return t.parent
}

func (t *Tree) SetParent(parent Node) {
	t.parent = parent
}

func (t *Tree) String() string {
	return String(t)
}
