package tree

import (
	"bytes"
	"fmt"
	"math"

	"github.com/andreasstrack/data"
)

type Node interface {
	GetValue() data.Value

	GetChildren() []Node
	Add(child Node) error
	Insert(child Node, index int) error
	Remove(index int) error
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

func GetAllNodesOfTree(root Node) []Node {
	allNodes := make([]Node, 0)
	queue := data.NewFifoQueue()
	queue.Insert(root)
	for !queue.IsEmpty() {
		cur := queue.Pop().(Node)
		allNodes = append(allNodes, cur)
		for i := 0; i < len(cur.GetChildren()); i++ {
			queue.Insert(cur.GetChildren()[i])
		}
	}
	return allNodes
}
