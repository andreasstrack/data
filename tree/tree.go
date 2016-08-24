package tree

import "github.com/andreasstrack/a-ds/datastructures"

// Parent represents a parent, which can have
// children.
type Parent interface {
	GetChildren() []Child
	Add(c Child)
	Insert(c Child, index int)
	Remove(index int)
}

// A Child has a parent.
type Child interface {
	GetParent() Parent
	SetParent(p Parent)
}

// A Node has a parent, children, and a value.
type Node struct {
	parent   Parent
	children []Child
	datastructures.Value
}

func (n *Node) GetChildren() []Child {
	return n.children
}

func (n *Node) Add(c Child) {
	n.children = append(n.children, c)
	c.SetParent(n)
}

func (n *Node) Insert(c Child, index int) {
	n.children = append(append(n.children[:index], c), n.children[index:]...)
}

func (n *Node) Remove(index int) {
	n.children = append(n.children[:index], n.children[index+1:]...)
}

func (n *Node) GetParent() Parent {
	return n.parent
}

func (n *Node) SetParent(p Parent) {
	n.parent = p
}

// NewTree will generate a new tree with only the
// root node containing v as its value.
func NewTree(v datastructures.Value) (root *Node) {
	n := Node{nil, make([]Child, 0), v}
	return &n
}
