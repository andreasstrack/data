package datastructures

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
