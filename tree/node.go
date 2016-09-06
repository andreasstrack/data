package tree

import (
	"bytes"
	"fmt"
	"math"

	"github.com/andreasstrack/datastructures"
)

type Node interface {
	GetValue() datastructures.Value

	GetChildren() []Node
	Add(child Node)
	Insert(child Node, index int) error
	Remove(index int) error
	GetParent() Node
	SetParent(n Node)
}

func String(n Node) string {
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%s", n.GetValue().String()))
	children := n.GetChildren()
	if len(children) > 0 {
		b.WriteString(" -> [")
		for i := range children {
			b.WriteString(String(children[i]))
			if i < len(children)-1 {
				b.WriteString(" ")
			}
		}
		b.WriteString("]")
	}
	// b.WriteString("]")
	return b.String()
}

func Size(n Node) int {
	if n == nil {
		return 0
	}
	size := 1
	children := n.GetChildren()
	for i := range children {
		size = size + Size(children[i])
	}
	return size
}

func Depth(n Node) int {
	if n == nil {
		return 0
	}
	children := n.GetChildren()
	maxChildDepth := 0
	for i := range children {
		maxChildDepth = int(math.Max(float64(maxChildDepth), float64(Depth(children[i]))))
	}
	return maxChildDepth + 1
}

func BranchingFactor(n Node) int {
	if n == nil {
		return 0
	}
	children := n.GetChildren()
	bf := len(children)
	for i := range children {
		bf = int(math.Max(float64(bf), float64(BranchingFactor(children[i]))))
	}
	return bf
}

func GetParentAndUncles(n Node) (uncles []Node) {
	if n == nil {
		return
	}
	parent := n.GetParent()
	if parent == nil {
		return
	}
	grandparent := parent.GetParent()
	if grandparent == nil {
		uncles = append(uncles, parent)
		return
	}
	return grandparent.GetChildren()
}

func GetSelfAndSiblings(n Node) (siblings []Node) {
	if n == nil {
		return
	}
	parent := n.GetParent()
	if parent == nil {
		siblings = append(siblings, n)
		return
	}
	siblings = parent.GetChildren()
	return
}

func GetSelfSiblingsAndCousins(n Node) (cousins []Node) {
	uncles := GetParentAndUncles(n)
	if len(uncles) == 0 {
		cousins = append(cousins, n)
		return
	}
	for i := range uncles {
		cousins = append(cousins, uncles[i].GetChildren()...)
	}
	return
}

func GetChildrenAndNephews(n Node) (nephews []Node) {
	siblings := GetSelfAndSiblings(n)
	for i := range siblings {
		nephews = append(nephews, siblings[i].GetChildren()...)
	}
	return
}

func GetRoot(n Node) (root Node) {
	if n == nil {
		return
	}
	var child, parent Node
	child = n
	for root == nil {
		parent = child.GetParent()
		if parent == nil {
			root = child
		}
		child = parent
		parent = nil
	}

	return root
}

func GetAllNodesOfTree(n Node) (nodes []Node) {
	root := GetRoot(n)
	if root == nil {
		return
	}
	ni := NewNodeIterator(root, newDefaultChildIterator, DepthFirst)
	for ni.HasNext() {
		nodes = append(nodes, ni.Next().(Node))
	}
	return
}
